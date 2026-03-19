package server

import (
	"context"
	"pulse/utils/config"
	interceptor "pulse/utils/middleware"

	"net"
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServerParams struct {
	fx.In
	Logger      *zap.Logger
	Interceptor *interceptor.GRPCAuthInterceptor
}

func NewGRPCServer(p GRPCServerParams) *grpc.Server {
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(p.Interceptor.Unary(p.Logger)),
		grpc.MaxRecvMsgSize(1024*1024*100),
		grpc.MaxSendMsgSize(1024*1024*100),
	)
	return grpcServer
}

func RunGRPCServer(lc fx.Lifecycle, grpcServer *grpc.Server, env *config.Env, logger *zap.Logger) {
	reflection.Register(grpcServer)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", env.Addr)
			if err != nil {
				logger.Error(err.Error())
				return err
			}
			go func() {
				err := grpcServer.Serve(ln)
				if err != nil {
					logger.Error(err.Error())
				}
			}()
			go func() {
				grpcWebServer := grpcweb.WrapServer(grpcServer,
					grpcweb.WithOriginFunc(func(origin string) bool { return true }),
				)

				httpServer := http.Server{
					Addr: ":80",
					Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						if grpcWebServer.IsGrpcWebRequest(r) || grpcWebServer.IsAcceptableGrpcCorsRequest(r) {
							grpcWebServer.ServeHTTP(w, r)
							return
						}

						switch r.URL.Path {
						case "/health":
							w.WriteHeader(http.StatusOK)
							w.Write([]byte("Success"))
						default:
							w.WriteHeader(http.StatusOK)
							w.Write([]byte("Success"))
						}
					}),
				}
				if err := httpServer.ListenAndServe(); err != nil {
					logger.Error(err.Error())
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			grpcServer.GracefulStop()
			return nil
		},
	})
}
