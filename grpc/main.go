package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"m/v2/pb"
	"m/v2/sever"
)

func main() {

	s := grpc.NewServer()
	todo := sever.NewTodoService()
	reflection.Register(s)

	pb.RegisterToDoServiceServer(s, todo)

	tl, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(fmt.Println("Error starting tcp listener on port 8765", err))
	}
	s.Serve(tl)
}
