package middleware

import (
	"context"
	"fmt"
	"pulse/utils/config"
	"pulse/utils/jwt"
	"pulse/utils/other"
	"strings"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type GRPCAuthInterceptor struct {
	config                *config.Env
	verifyJwtTokenManager *jwt.VerifyJwtTokenManager
	utils                 *other.Utils
}

func NewGRPCAuthInterceptor(config *config.Env, verifyJwtTokenManager *jwt.VerifyJwtTokenManager, utils *other.Utils) *GRPCAuthInterceptor {
	return &GRPCAuthInterceptor{
		config:                config,
		verifyJwtTokenManager: verifyJwtTokenManager,
		utils:                 utils,
	}
}

func (i *GRPCAuthInterceptor) Unary(logger *zap.Logger) grpc.UnaryServerInterceptor {

	publicMethods := map[string]bool{
		// admin auth
		// "/userProto.UserAuthService/UserSignUp":          true,
	}

	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp any, err error) {

		defer func() {
			if r := recover(); r != nil {
				logger.Error("panic in gRPC handler",
					zap.Any("error", r),
				)
				err = i.utils.GrpcErr(codes.Internal, "Something went wrong!")
			}
		}()

		if publicMethods[info.FullMethod] {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
		}

		authHeader := md.Get("authorization")
		if len(authHeader) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "authorization header not found")
		}

		tokenStr := strings.TrimPrefix(authHeader[0], "Bearer ")
		if tokenStr == "" {
			return nil, status.Errorf(codes.Unauthenticated, "invalid token format")
		}

		_, err = i.verifyJwtTokenManager.VerifyToken(tokenStr)
		if err != nil {
			fmt.Println("err", err)
			return nil, status.Errorf(codes.Unauthenticated, "invalid or expired token")
		}

		return handler(ctx, req)
	}
}
