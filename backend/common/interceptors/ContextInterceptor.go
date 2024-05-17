package interceptors

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func ContextInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// resource
	var lang string
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		locales := md["resource"]
		lang = locales[0]
	} else {
		lang = "en"
	}
	newCtx := context.WithValue(ctx, "resource", lang)
	resp, err := handler(newCtx, req)
	return resp, err
}
