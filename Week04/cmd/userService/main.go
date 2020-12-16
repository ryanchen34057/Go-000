package main

import (
	"Go-000/Week04/api/user"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	signals := make(chan os.Signal)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go processSignals(signals)

	userService, err := InitializeUserService()
	if err != nil {
		fmt.Printf("Failed to start service: %s\n", err)
		os.Exit(2)
	}
	grpcServer := grpc.NewServer()
	userServiceServer := user.NewUserServer(&userService)
	user.RegisterUserServiceServer(grpcServer, userServiceServer)
	grpcServer.Serve(lis)
}

func processSignals(sigs chan os.Signal) {
	<-sigs
	os.Exit(2)
}
