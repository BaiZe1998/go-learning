package server

import (
	v1 "blog-service/api/blog/v1"
	"blog-service/internal/conf"
	"blog-service/internal/service"

	nethttp "net/http"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, blog *service.BlogService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
		// 添加CORS支持
		http.Filter(
			func(handler nethttp.Handler) nethttp.Handler {
				return nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
					w.Header().Set("Access-Control-Allow-Origin", "*")
					w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
					w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

					if r.Method == "OPTIONS" {
						w.WriteHeader(200)
						return
					}

					handler.ServeHTTP(w, r)
				})
			},
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
	v1.RegisterBlogHTTPServer(srv, blog)
	return srv
}
