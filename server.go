package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"net/http"
	"ttt/helper"
	"ttt/services"
)

func main(){

	//启动服务端

	doubleTls()
	//tlsGrpc()


}


func tlsGrpc(){

	file, err2 := credentials.NewServerTLSFromFile("keys/server.crt", "keys/server.key")

	if err2!=nil{
		log.Fatal(err2)
	}

	rpcServer := grpc.NewServer(grpc.Creds(file))

	services.RegisterProdServiceServer(rpcServer,new(services.ProdService))
	lis,_ := net.Listen("tcp",":8081")
	rpcServer.Serve(lis)

}


func rawCon(){
	//---
	rpcServer := grpc.NewServer()
	services.RegisterProdServiceServer(rpcServer,new(services.ProdService))
	lis,_ := net.Listen("tcp",":9091")
	rpcServer.Serve(lis)
	//----
}

func httpCon(){
	//---
	rpcServer := grpc.NewServer()
	services.RegisterProdServiceServer(rpcServer,new(services.ProdService))


		 mux:=http.NewServeMux()

		 mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

			fmt.Println(request.Proto)
			fmt.Println(request.Header)
			fmt.Println(request.Body)
			fmt.Println("-----",request)
			 rpcServer.ServeHTTP(writer,request)


		 })

		server := http.Server{
			Addr: ":8081",
			Handler: mux,
		}

		 server.ListenAndServeTLS("../keys/server.crt","../keys/server.key")

	//---


}


func doubleTls(){
	creds :=helper.GetDoubleTls()
	rpcServer := grpc.NewServer(grpc.Creds(creds))
	services.RegisterProdServiceServer(rpcServer,new(services.ProdService))//商品服务

	services.RegisterOrderServiceServer(rpcServer,new(services.OrderService))//订单服务
	lis,_ := net.Listen("tcp",":8081")
	fmt.Println("start listening")
	rpcServer.Serve(lis)
}