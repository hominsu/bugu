package server

import (
	"context"

	buguV1 "bugu/api/bugu/service/v1"
	"bugu/app/bugu/service/internal/conf"
	"bugu/app/bugu/service/internal/pkg/middleware/auth"
	"bugu/app/bugu/service/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
)

func NewSkipRoutersMatcher() selector.MatchFunc {
	skipList := make(map[string]struct{})
	skipList["/bugu.service.v1.Bugu/Login"] = struct{}{}
	skipList["/bugu.service.v1.Bugu/Register"] = struct{}{}

	return func(ctx context.Context, operation string) bool {
		if _, ok := skipList[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, sc *conf.Jwt, is *service.BuguService, logger log.Logger) *http.Server {
	opts := []http.ServerOption{
		http.Middleware(
			recovery.Recovery(
				recovery.WithLogger(logger),
			),
			logging.Server(logger),
			ratelimit.Server(),
			selector.Server(
				auth.JwtAuth(sc.GetSecret()),
			).Match(NewSkipRoutersMatcher()).
				Build(),
			validate.Validator(),
		),
		http.Filter(
			handlers.CORS(
				handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
				handlers.AllowedOrigins([]string{"*"}),
			),
		),
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
	srv := http.NewServer(opts...)

	buguV1.RegisterBuguHTTPServer(srv, is)
	return srv
}
