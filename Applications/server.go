package app

import (
	"fmt"
	"log"
	"net"
	"net/http"
	protobuf "practise/applications/protobuf"
	"time"

	"google.golang.org/grpc"
)

func StartGraphqlServer() {
	handler := Handler()
	http.Handle("/graphQl", handler)
	log.Println("App serving on graphql server over address: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func StartGrpcCalculatorServer() {

	listen, err := net.Listen("tcp4", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err.Error())
	}

	grpcService := NewGrpcApp()
	grpcServer := grpc.NewServer()
	protobuf.RegisterArithmeticServiceServer(grpcServer, grpcService)

	log.Println("Calculator App serving on gRPC server over address: grpc://localhost:9000")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC server over port 9000: %v", err)
	}
}

func StartRestCalculatorServer() {

	restApp := NewRestApp()
	http.HandleFunc("/GetAddition", restApp.GetAddition)
	http.HandleFunc("/GetSubtraction", restApp.GetSubtraction)
	http.HandleFunc("/GetMultiplication", restApp.GetMultiplication)
	http.HandleFunc("/GetDivision", restApp.GetDivision)

	log.Println("Calculator App serving on Rest server over address: http://localhost:8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalf("failed to serve gRPC server over port 8000: %v", err)
	}
}

func PingOne(domain string, port string) {

	address := domain + ":" + port
	timeout := time.Duration(2 * time.Second)

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		fmt.Printf("[DOWN] %v is unreachable, \nError: %v", domain, err)
	} else {
		fmt.Printf("[UP] %v is reachable, \nFrom: %v\nTo: %v", domain, conn.LocalAddr(), conn.RemoteAddr())
	}
}
