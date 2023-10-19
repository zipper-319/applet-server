package apiHook

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	httpKratos "github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/proto"
	"net/http"
	"reflect"
	"runtime/debug"
	"strings"
	"time"
)

const (
	baseContentType = "application"
)

type LoggerKey struct{}

type Response struct {
	Code     int32  `json:"code"`
	Message  string `json:"message"`
	Reason   string `json:"reason"`
	MetaData any    `json:"metadata"`
}

func Hook(logger log.Logger) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			defer func() {
				if err := recover(); err != nil {
					log.NewHelper(logger).Error(err)
					fmt.Println(string(debug.Stack()))
				}
			}()
			var (
				kind      string
				operation string
			)
			code := 0
			message := "SUCCESS"
			startTime := time.Now()
			apiLogger := logger

			if info, ok := transport.FromServerContext(ctx); ok {
				kind = info.Kind().String()
				operation = info.Operation()
			}

			ctx = context.WithValue(ctx, LoggerKey{}, log.NewHelper(apiLogger))
			reply, err = handler(ctx, req)
			var result interface{}
			var version string
			if se := errors.FromError(err); se != nil {
				code = int(se.Code)
				message = se.Reason
			} else {
				rr := reflect.ValueOf(reply).Elem()
				if rr.Kind() != reflect.Invalid {
					result = rr.Interface()
				}
			}
			isTimeout := false
			latency := time.Since(startTime).Milliseconds()
			if latency > 300 {
				isTimeout = true
			}
			level, stack := extractError(err)
			if isTimeout {
				level = log.LevelWarn
			}
			log.WithContext(ctx, apiLogger).Log(level,
				"component", kind,
				"version", version,
				"operation", operation,
				"args", reflect.ValueOf(req).Elem().Interface(),
				"code", code,
				"message", message,
				"stack", stack,
				"result", result,
				"isTimeout", fmt.Sprintf("timeout is %t", isTimeout),
				"latency", latency)
			return
		}
	}
}

// extractError returns the string of the error
func extractError(err error) (log.Level, string) {
	if err != nil {
		return log.LevelError, fmt.Sprintf("%+v", err)
	}
	return log.LevelInfo, ""
}

// responseEncoder encodes the object to the HTTP response.
func ResponseEncoder(w http.ResponseWriter, r *http.Request, v interface{}) error {
	if v == nil {
		return nil
	}
	if rd, ok := v.(httpKratos.Redirector); ok {
		url, code := rd.Redirect()
		http.Redirect(w, r, url, code)
		return nil
	}

	if result, ok := v.(proto.Message); ok {

		codec, _ := codecForRequest(r, "Accept")
		respWrapper := Response{
			Code:     200,
			Message:  "success",
			MetaData: result,
		}
		data, err := codec.Marshal(respWrapper)
		if err != nil {
			return err
		}
		contentType := strings.Join([]string{baseContentType, codec.Name()}, "/")
		w.Header().Set("Content-Type", contentType)
		_, err = w.Write(data)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("result is not proto.message")

}

// errorEncoder encodes the error to the HTTP response.
func ErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	se := errors.FromError(err)
	codec, _ := codecForRequest(r, "Accept")
	body, err := codec.Marshal(se)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	contentType := strings.Join([]string{baseContentType, codec.Name()}, "/")
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(int(se.Code))
	_, _ = w.Write(body)
}

// codecForRequest get encoding.Codec via http.Request
func codecForRequest(r *http.Request, name string) (encoding.Codec, bool) {
	for _, accept := range r.Header[name] {
		codec := encoding.GetCodec(contentSubtype(accept))
		if codec != nil {
			return codec, true
		}
	}
	return encoding.GetCodec("json"), false
}

func contentSubtype(contentType string) string {
	left := strings.Index(contentType, "/")
	if left == -1 {
		return ""
	}
	right := strings.Index(contentType, ";")
	if right == -1 {
		right = len(contentType)
	}
	if right < left {
		return ""
	}
	return contentType[left+1 : right]
}

func GetLoggerFromContext(ctx context.Context) *log.Helper {
	if logger, ok := ctx.Value(LoggerKey{}).(*log.Helper); ok {
		return logger
	} else {
		return log.NewHelper(log.DefaultLogger)
	}
}
