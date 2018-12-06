package main

import (
	"context"
	pb "github.com/brymck/opto-demo/src/frontend/genproto"
)

func (fe *frontendServer) sayHello(ctx context.Context, name string) (string, error) {
	resp, err := pb.NewHelloServiceClient(fe.helloSvcConn).GetGreeting(ctx, &pb.GreetingRequest{Name: name})
	return resp.GetMessage(), err
}
