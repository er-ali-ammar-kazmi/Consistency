package app

import (
	"context"
	"fmt"

	protobuf "practise/applications/protobuf"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GrpcApp implements the GRPC interface
type GrpcApp struct {
	calci Calculator
	protobuf.UnimplementedArithmeticServiceServer
}

// NewGrpcApp creates a new GrpcApp
func NewGrpcApp() *GrpcApp {
	return &GrpcApp{calci: NewCalculator()}
}

// GetAddition gets the result of adding Operators a and b
func (grpcService GrpcApp) GetAddition(ctx context.Context, req *protobuf.Operators) (*protobuf.Msg, error) {
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
func (grpcService GrpcApp) GetSubtraction(ctx context.Context, req *protobuf.Operators) (*protobuf.Msg, error) {

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
func (grpcService GrpcApp) GetMultiplication(ctx context.Context, req *protobuf.Operators) (*protobuf.Msg, error) {

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
func (grpcService GrpcApp) GetDivision(ctx context.Context, req *protobuf.Operators) (*protobuf.Msg, error) {

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
