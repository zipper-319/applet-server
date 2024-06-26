package server

import (
	"applet-server/api/v2/applet"
	"applet-server/internal/conf"
	"applet-server/internal/pkg/apiHook"
	jwtUtil "applet-server/internal/pkg/jwt"
	"applet-server/internal/pkg/log"
	"applet-server/internal/service"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, middlewares http.ServerOption, vdSer *service.VoiceDataOperationService, account *service.AccountService,
	speakerService *service.CloneSpeakerService, ttsService *service.TTSServiceService, chat *service.ChatService, feedbackService *service.FeedbackService,
	logger *log.MyLogger) *http.Server {
	var opts = []http.ServerOption{
		middlewares,
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	opts = append(opts, http.ResponseEncoder(apiHook.ResponseEncoder))
	opts = append(opts, http.ErrorEncoder(apiHook.ErrorEncoder))
	srv := http.NewServer(opts...)

	swagger := openapiv2.NewHandler()
	srv.HandlePrefix("/q", swagger)

	applet.RegisterVoiceDataOperationHTTPServer(srv, vdSer)
	applet.RegisterAccountHTTPServer(srv, account)
	applet.RegisterCloneSpeakerHTTPServer(srv, speakerService)
	applet.RegisterTTSServiceHTTPServer(srv, ttsService)
	applet.RegisterFeedbackHTTPServer(srv, feedbackService)
	RegisterFormDataHandler(srv, vdSer)
	RegisterChatWebsocketServer(srv, chat, logger)
	return srv
}

func NewMiddlewares(logger *log.MyLogger, appConfig *conf.App) http.ServerOption {
	return http.Middleware(
		validate.Validator(),
		tracing.Server(),
		metrics.Server(),
		jwtUtil.Server(logger, appConfig.Auth.Key, appConfig.Auth.Expire.AsDuration()),
		apiHook.Hook(logger),
	)
}
