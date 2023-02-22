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
	err2 := godotenv.Load()
	if err2 != nil {
		log.Fatalf("Error loading .env file %v\n", err2)
	}

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
