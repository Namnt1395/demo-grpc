package main

import (
	grpc_server "client_grpc/server"
	pub "server-grpc/pub/pubsub"
	"fmt"
	"golang.org/x/net/context"
	"io"
	"log"
)

func main() {
	//Connect grpc server
	client := grpc_server.ConnectGRPC()
	stream, err := client.Subscribe(
		context.Background(), &pub.String{Key: "data1,data2"},
	)
	if err != nil {
		log.Fatal(err)
	}
	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		switch reply.GetKey() {
		case "SSP":
			fmt.Println("SSP: " + reply.GetValue())
			break
		case "DSP":
			fmt.Println("DSP: " + reply.GetValue())
			break
		default:
			fmt.Println("DEFAULT: " + reply.GetValue())
			break

		}
	}
}
