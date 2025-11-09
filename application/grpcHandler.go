package app

import (
	"context"
	"fmt"

	protobuf "practise/application/protobuf"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcApp struct {
	calci Calculator
	protobuf.UnimplementedArithmeticServiceServer
}

func NewGrpcApp() *GrpcApp {
	return &GrpcApp{calci: NewCalculator()}
}

func (grpcService GrpcApp) GetAddition(ctx context.Context, req *protobuf.Operators) (*protobuf.Msg, error) {

	res := &protobuf.Msg{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return res, status.Error(codes.InvalidArgument, "missing required argument!")
	}

	output := grpcService.calci.Addition(req.A, req.B)
	result := fmt.Sprintf("Addition of %f and %f is %f", req.A, req.B, output)

	(*res).Value = result
	return res, nil
}

func (grpcService GrpcApp) GetSubtraction(ctx context.Context, req *protobuf.Operators) (*protobuf.Msg, error) {

	res := &protobuf.Msg{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return res, status.Error(codes.InvalidArgument, "missing required argument!")
	}

	output := grpcService.calci.Subtraction(req.A, req.B)
	result := fmt.Sprintf("Subtraction of %f and %f is %f", req.A, req.B, output)

	(*res).Value = result
	return res, nil
}

func (grpcService GrpcApp) GetMultiplication(ctx context.Context, req *protobuf.Operators) (*protobuf.Msg, error) {

	res := &protobuf.Msg{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return res, status.Error(codes.InvalidArgument, "missing required argument!")
	}

	output := grpcService.calci.Multiplication(req.A, req.B)
	result := fmt.Sprintf("Multiplication of %f and %f is %f", req.A, req.B, output)

	(*res).Value = result
	return res, nil
}

func (grpcService GrpcApp) GetDivision(ctx context.Context, req *protobuf.Operators) (*protobuf.Msg, error) {

	res := &protobuf.Msg{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return res, status.Error(codes.InvalidArgument, "missing required argument!")
	}

	output := grpcService.calci.Division(req.A, req.B)
	result := fmt.Sprintf("Division of %f and %f is %f", req.A, req.B, output)

	(*res).Value = result
	return res, nil
}
