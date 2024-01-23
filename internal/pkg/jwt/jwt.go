package jwtUtil

import (
	"applet-server/internal/pkg/log"
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"strings"
	"time"
)

type Identifier struct{}

type Option func(*options)

var NotVerify = []string{
	"login",
	"GetTtsConfig",
	"GetVersion",
	"GetSpeaker",
	"UploadFiles",
	"chat",
}

const (

	// authorizationKey holds the key used to store the JWT Token in the request tokenHeader.
	AuthorizationKey string = "Authorization"

	// reason holds the error reason.
	reason string = "UNAUTHORIZED"
)

var (
	ErrMissingJwtToken        = errors.Unauthorized(reason, "jwt token is missing")
	ErrMissingKey             = errors.Unauthorized(reason, "jwt key is missing")
	ErrTokenInvalid           = errors.Unauthorized(reason, "Token is invalid")
	ErrTokenExpired           = errors.Unauthorized(reason, "JWT token has expired")
	ErrTokenParseFail         = errors.Unauthorized(reason, "Fail to parse JWT token ")
	ErrUnSupportSigningMethod = errors.Unauthorized(reason, "Wrong signing method")
	ErrWrongContext           = errors.Unauthorized(reason, "Wrong context for middleware")
	ErrNeedTokenProvider      = errors.Unauthorized(reason, "Token provider is missing")
	ErrSignToken              = errors.Unauthorized(reason, "Can not sign token.Is the key correct?")
	ErrGetKey                 = errors.Unauthorized(reason, "Can not get key while signing token")
)

var keyFunc jwt.Keyfunc
var expireTime time.Duration

func init() {
	keyFunc = KeyProvider("")
	expireTime = time.Duration(30)
}

// Parser is a jwt parser
type options struct {
	signingMethod jwt.SigningMethod
	claims        jwt.Claims
	tokenHeader   map[string]interface{}
}

// WithClaims with customer claim
// If you use it in Server, f needs to return a new jwt.Claims object each time to avoid concurrent write problems
// If you use it in Client, f only needs to return a single object to provide performance
func WithClaims(claims jwt.Claims) Option {
	return func(o *options) {
		o.claims = claims
	}
}

func Server(logger *log.MyLogger, jwtKey string, expire time.Duration, opts ...Option) middleware.Middleware {
	keyFunc = KeyProvider(jwtKey)
	expireTime = expire
	o := &options{
		signingMethod: jwt.SigningMethodHS256,
		claims:        &IdentityClaims{},
	}
	for _, opt := range opts {
		opt(o)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if header, ok := transport.FromServerContext(ctx); ok {

				isNotVerify := false
				for _, operation := range NotVerify {
					if strings.Contains(header.Operation(), operation) {
						isNotVerify = true
						break
					}
				}
				if !isNotVerify {
					jwtToken := header.RequestHeader().Get(AuthorizationKey)
					logger.Infof("jwtToken:%s", jwtToken)
					if jwtToken == "" {
						return nil, ErrMissingJwtToken
					}

					var (
						tokenInfo *jwt.Token
						err       error
					)

					if o.claims != nil {
						tokenInfo, err = jwt.ParseWithClaims(jwtToken, o.claims, keyFunc)
					} else {
						tokenInfo, err = jwt.Parse(jwtToken, keyFunc)
					}
					if err != nil {
						ve, ok := err.(*jwt.ValidationError)
						if !ok {
							return nil, errors.Unauthorized(reason, err.Error())
						}
						if ve.Errors&jwt.ValidationErrorMalformed != 0 {
							return nil, ErrTokenInvalid
						}
						if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
							return nil, ErrTokenExpired
						}
						if ve.Inner != nil {
							return nil, ve.Inner
						}
						return nil, ErrTokenParseFail
					}
					if !tokenInfo.Valid {
						return nil, ErrTokenInvalid
					}
					if tokenInfo.Method != o.signingMethod {
						return nil, ErrUnSupportSigningMethod
					}
					if claims, ok := tokenInfo.Claims.(*IdentityClaims); ok {
						ctx = context.WithValue(ctx, Identifier{}, claims)
						logger.Infof("tokenInfo.Claims is %#v", claims)
					} else {
						logger.Error("tokenInfo.Claims is not *IdentityClaims")
					}

				}
				return handler(ctx, req)
			}
			return nil, ErrWrongContext
		}

	}
}

func GenerateJwtToken(kf jwt.Keyfunc, opts ...Option) (string, error) {
	o := &options{
		signingMethod: jwt.SigningMethodHS256,
	}
	for _, opt := range opts {
		opt(o)
	}
	token := jwt.NewWithClaims(o.signingMethod, o.claims)
	if o.tokenHeader != nil {
		for k, v := range o.tokenHeader {
			token.Header[k] = v
		}
	}
	key, err := kf(token)
	if err != nil {
		return "", err
	}

	if tokenStr, err := token.SignedString(key); err != nil {
		return "", err
	} else {
		return tokenStr, nil
	}
}

func KeyProvider(key string) jwt.Keyfunc {
	return func(*jwt.Token) (interface{}, error) {
		return []byte(key), nil
	}
}

type IdentityClaims struct {
	Username    string `json:"username"`
	PhoneNumber string `json:"phoneNumber"`
	Role        int    `json:"role"`
	SessionID   string `json:"session"`
	jwt.RegisteredClaims
}

func GetToken(username, PhoneNumber string, role int) (string, error) {
	now := time.Now()
	id, _ := uuid.NewUUID()
	claims := IdentityClaims{
		Username:    username,
		PhoneNumber: PhoneNumber,
		Role:        role,
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  []string{username},
			ExpiresAt: jwt.NewNumericDate(now.Add(expireTime * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
			ID:        id.String(),
		},
	}
	return GenerateJwtToken(keyFunc, WithClaims(claims))
}

func ParseToken(token, key string) (*IdentityClaims, error) {
	identifier := IdentityClaims{}
	tokenInfo, err := jwt.ParseWithClaims(token, &identifier, KeyProvider(key))
	if err != nil {
		return nil, err
	}

	if claims, ok := tokenInfo.Claims.(*IdentityClaims); ok {
		log.Debug("tokenInfo  userName:", claims.Username)
		return claims, nil
	}
	return nil, err
}

func GetTokenInfo(ctx context.Context) (tokenInfo *IdentityClaims, ok bool) {
	tokenInfo, ok = ctx.Value(Identifier{}).(*IdentityClaims)
	return
}
