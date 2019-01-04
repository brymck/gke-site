package main

import (
	"context"
	pb "github.com/brymck/gke-site/src/frontend/server/genproto"
)

func (fe *frontendServer) sayHello(ctx context.Context, name string) (*pb.Greeting, error) {
	resp, err := pb.NewHelloServiceClient(fe.helloSvcConn).GetGreeting(ctx, &pb.GreetingRequest{Name: name})
	return resp, err
}
