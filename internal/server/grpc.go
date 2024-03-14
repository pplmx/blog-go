package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/pplmx/blog-go/internal/service"
	blogv1 "github.com/pplmx/pb/blog/v1"
	"github.com/spf13/viper"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	logger log.Logger,
	categoryService *service.CategoryService,
	commentService *service.CommentService,
	postService *service.PostService,
	tagService *service.TagService,
) *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
			validate.Validator(),
		),
	}

	addr := viper.GetString("grpc.addr")
	if addr != "" {
		opts = append(opts, grpc.Address(addr))
	}

	timeout := viper.GetDuration("grpc.timeout")
	if timeout > 0 {
		opts = append(opts, grpc.Timeout(timeout))
	}

	srv := grpc.NewServer(opts...)

	// Register the service to the server.
	blogv1.RegisterCategoryServiceServer(srv, categoryService)
	blogv1.RegisterCommentServiceServer(srv, commentService)
	blogv1.RegisterPostServiceServer(srv, postService)
	blogv1.RegisterTagServiceServer(srv, tagService)
	return srv
}
