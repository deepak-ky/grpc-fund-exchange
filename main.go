package main

import (
	"context"
	pb "grpc-fund-transfer/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type myGRPCServer struct {
	pb.UnimplementedTransferServer
}

func (m *myGRPCServer) Create(ctx context.Context,req *pb.CreateRequest) (*pb.CreateResponse, error) {
	log.Println("Create has been called")
	return &pb.CreateResponse{Pdf: []byte("TODO")}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer()
	myTransferServer := &myGRPCServer{}
	pb.RegisterTransferServer(s,myTransferServer)

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
