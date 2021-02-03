package main

import (
	"server-grpc/api"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	groupEvent := router.PathPrefix("/api/push/").Subrouter()
	groupEvent.HandleFunc("/event", api.ApiPush)

	err := http.ListenAndServe(":8084", router)

	if err != nil {
		fmt.Println(err.Error())
	}
}
