package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	pb "github.com/MarioScripts/vlr-api/proto/gen/vlr/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	pb.VlrServer
}

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:50051", "gRPC server endpoint")
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	godotenv.Load()

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	addr := fmt.Sprintf("%s:%s", host, port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen of: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterVlrServer(s, &Server{})

	go func() {
		log.Fatalln(s.Serve(lis))
	}()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	e := pb.RegisterVlrHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if e != nil {
		log.Fatalf("%v", e)
	}

	http.ListenAndServe(":8081", mux)

}
