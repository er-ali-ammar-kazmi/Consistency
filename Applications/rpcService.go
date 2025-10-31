package app

import (
	"context"
	"fmt"
	"log"
	"net"

	protobuf "practise/applications/protobuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Adapter implements the GRPCPort interface
type Adapter struct {
	calci Calculator
}

// NewAdapter creates a new Adapter
func NewAdapter(app Calculator) *Adapter {
	return &Adapter{calci: app}
}

// GetAddition gets the result of adding Operators a and b
func (grpcService Adapter) GetAddition(ctx context.Context, req *protobuf.Operators) (*protobuf.Msg, error) {
	ans := &protobuf.Msg{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	output := grpcService.calci.Addition(req.A, req.B)
	result := fmt.Sprintf("Addition of %f and %f is %f", req.A, req.B, output)
	ans = &protobuf.Msg{
		Value: result,
	}

	return ans, nil
}

// GetSubtraction gets the result of subtracting Operators a and b
func (grpcService Adapter) GetSubtraction(ctx context.Context, req *protobuf.Operators) (*protobuf.Msg, error) {

	ans := &protobuf.Msg{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	output := grpcService.calci.Subtraction(req.A, req.B)
	result := fmt.Sprintf("Subtraction of %f and %f is %f", req.A, req.B, output)
	ans = &protobuf.Msg{
		Value: result,
	}

	return ans, nil
}

// GetMultiplication gets the result of multiplying Operators a and b
func (grpcService Adapter) GetMultiplication(ctx context.Context, req *protobuf.Operators) (*protobuf.Msg, error) {

	ans := &protobuf.Msg{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	output := grpcService.calci.Multiplication(req.A, req.B)
	result := fmt.Sprintf("Multiplication of %f and %f is %f", req.A, req.B, output)
	ans = &protobuf.Msg{
		Value: result,
	}

	return ans, nil
}

// GetDivision gets the result of dividing Operators a and b
func (grpcService Adapter) GetDivision(ctx context.Context, req *protobuf.Operators) (*protobuf.Msg, error) {

	ans := &protobuf.Msg{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	output := grpcService.calci.Division(req.A, req.B)
	result := fmt.Sprintf("Division of %f and %f is %f", req.A, req.B, output)
	ans = &protobuf.Msg{
		Value: result,
	}

	return ans, nil
}

// Run registers the ArithmeticServiceServer to a grpcServer and serves on
// the specified port
func (grpcService Adapter) Run() {

	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}

	grpcServer := grpc.NewServer()
	protobuf.RegisterArithmeticServiceServer(grpcServer, grpcService)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC server over port 9000: %v", err)
	}
}
