package main

import (
	"context"
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
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type Server struct {
	pb.VlrServer
	grpc_health_v1.HealthServer
}

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
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())

	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	mux := runtime.NewServeMux(runtime.WithHealthzEndpoint(grpc_health_v1.NewHealthClient(conn)))
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	e := pb.RegisterVlrHandlerFromEndpoint(ctx, mux, addr, opts)
	if e != nil {
		log.Fatalf("%v", e)
	}
	http.ListenAndServe(":8081", mux)

}
