package main

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/MarioScripts/vlr-api/proto/gen/vlr/v1"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type Server struct {
	pb.VlrServer
}

func main() {
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

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v\n", err)
	}

}
