package main

import (
	pub "server-grpc/pub/pubsub"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main()  {
	// khởi tạo một đối tượng gRPC service
	grpcServer := grpc.NewServer()
	s := pub.NewPubsubService()
	// đăng ký service với grpcServer (của gRPC plugin)
	pub.RegisterPubsubServiceServer(grpcServer, s)

	// cung cấp gRPC service trên port `1234`
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
