package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/pplmx/blog-go/internal/service"
	blogv1 "github.com/pplmx/pb/blog/v1"
	"github.com/spf13/viper"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	logger log.Logger,
	categoryService *service.CategoryService,
	commentService *service.CommentService,
	postService *service.PostService,
	tagService *service.TagService,
) *http.Server {
	opts := []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
			validate.Validator(),
		),
	}

	addr := viper.GetString("http.addr")
	if addr != "" {
		opts = append(opts, http.Address(addr))
	}

	timeout := viper.GetDuration("http.timeout")
	if timeout > 0 {
		opts = append(opts, http.Timeout(timeout))
	}

	srv := http.NewServer(opts...)

	// Register the service to the server.
	blogv1.RegisterCategoryServiceHTTPServer(srv, categoryService)
	blogv1.RegisterCommentServiceHTTPServer(srv, commentService)
	blogv1.RegisterPostServiceHTTPServer(srv, postService)
	blogv1.RegisterTagServiceHTTPServer(srv, tagService)
	return srv
}
