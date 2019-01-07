package main

import (
	"context"
	pb "github.com/brymck/gke-site/src/frontend/server/genproto"
)

func (fe *frontendServer) count(ctx context.Context) (*pb.CountResponse, error) {
	resp, err := pb.NewCountServiceClient(fe.countSvcConn).GetCount(ctx, &pb.CountRequest{})
	return resp, err
}

func (fe *frontendServer) sayHello(ctx context.Context, name string) (*pb.Greeting, error) {
	resp, err := pb.NewHelloServiceClient(fe.helloSvcConn).GetGreeting(ctx, &pb.GreetingRequest{Name: name})
	return resp, err
}

func (fe *frontendServer) getSquare(ctx context.Context, number float32) (*pb.SquareResponse, error) {
	resp, err := pb.NewSquareServiceClient(fe.squareSvcConn).GetSquare(ctx, &pb.SquareRequest{Number: number})
	return resp, err
}
