package server

import (
	pub "server-grpc/pub/pubsub"
	"google.golang.org/grpc"
	"log"
)

var gRPC *grpc.ClientConn

func ConnectGRPC() pub.PubsubServiceClient {
	var err error
	gRPC, err = grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := pub.NewPubsubServiceClient(gRPC)
	return client
}
func CloseGRPC() {
	_ = gRPC.Close()
}
