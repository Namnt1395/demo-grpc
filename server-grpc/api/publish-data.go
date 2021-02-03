package api

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"
	pub "server-grpc/pub/pubsub"
)

func ApiPush(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	data := queries.Get("data")
	key := queries.Get("key")
	client := connectGrpcServer()
	_, err := client.Publish(
		context.Background(), &pub.String{Key: key, Value: data},
	)
	if err != nil {
		log.Fatal(err)
	}
}
func connectGrpcServer() pub.PubsubServiceClient {
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := pub.NewPubsubServiceClient(conn)
	return client
}
