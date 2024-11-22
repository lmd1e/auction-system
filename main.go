package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"auction-system/api/grpc"
	"auction-system/infrastructure/data_access"
	"auction-system/workers"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	userRepo := data_access.NewPostgresUserRepository(db)
	lotRepo := data_access.NewPostgresLotRepository(db)
	bidRepo := data_access.NewPostgresBidRepository(db)
	auctionRepo := data_access.NewPostgresAuctionRepository(db)

	grpcServer := grpc.NewServer()
	grpc.RegisterAuctionServiceServer(grpcServer, grpc.NewAuctionService(userRepo, lotRepo, bidRepo, auctionRepo))

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		log.Println("Starting gRPC server on port 50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = grpc.RegisterAuctionServiceHandlerFromEndpoint(context.Background(), mux, "localhost:50051", opts)
	if err != nil {
		log.Fatalf("failed to register gateway: %v", err)
	}

	go func() {
		log.Println("Starting REST server on port 8080")
		if err := http.ListenAndServe(":8080", mux); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	worker := workers.NewAuctionWorker(auctionRepo, bidRepo, userRepo)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go worker.Start(ctx)

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down servers...")
	grpcServer.GracefulStop()
	cancel()
	log.Println("Servers gracefully stopped")
}
