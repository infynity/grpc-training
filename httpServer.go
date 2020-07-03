package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"ttt/helper"
	"ttt/services"
)

func main() {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(helper.GetDoubleTlsForClient())}
	err2 := services.RegisterProdServiceHandlerFromEndpoint(context.Background(), mux, "localhost:8081", opts) //8081  is grpc addr

	if err2 != nil {
		log.Fatal(err2)
	}

	hsv := http.Server{
		Addr: ":9988",
		Handler: mux,
	}

	hsv.ListenAndServe()
}
