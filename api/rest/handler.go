package rest

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func RegisterHandlers(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return auction.RegisterAuctionServiceHandler(ctx, mux, conn)
}

func StartRESTServer(ctx context.Context, grpcEndpoint string) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := auction.RegisterAuctionServiceHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts); err != nil {
		return err
	}

	return http.ListenAndServe(":8080", mux)
}
