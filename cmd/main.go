package main

import (
	"fmt"
	"github.com/zhulinwei/gin-demo/pkg/config"
	"github.com/zhulinwei/gin-demo/pkg/model/protobuf"
	"github.com/zhulinwei/gin-demo/pkg/router"
	"github.com/zhulinwei/gin-demo/pkg/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func GRPCRun(port string) {
	fmt.Printf("grpc run %v", port)
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	protobuf.RegisterGreeterServer(server, &service.GreeterServer{})
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


func main() {
	httpPort := config.ServerConfig().HttpPort
	grpcPort := config.ServerConfig().GrpcPort

	fmt.Println(httpPort, grpcPort)

	go GRPCRun(grpcPort)
	if err := router.BuildRoute().Run(httpPort); err != nil {
		log.Fatalf("server run failed: %v", err)
	}
}
