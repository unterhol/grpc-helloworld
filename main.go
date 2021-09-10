package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	helloworldpb "github.com/unterhol/grpc-helloworld/proto/helloworld"
)

//
//  Sample grpc server
//
//  some command history...
//
//  go mod init github.com/unterhol/grpc-helloworld
//  go mod tidy
//  go mod vendor
//  go run main.go
//
// docker build . --tag docker-grpc-helloworld
// docker run --publish 8080:8080 docker-grpc-helloworld
//
// [Johns-MacBook-Pro :  ~/go/src/github.com/unterhol/grpc-helloworld ] - grpcurl -plaintext localhost:8080 list
// grpc.reflection.v1alpha.ServerReflection
// helloworld.Greeter
//
// [Johns-MacBook-Pro :  ~/go/src/github.com/unterhol/grpc-helloworld ] - grpcurl -plaintext localhost:8080 helloworld.Greeter/SayHello
// {
//   "message": " world"
// }
//
// [Johns-MacBook-Pro :  ~/go/src/github.com/unterhol/grpc-helloworld ] - grpcurl -plaintext -d '{"name": "hello"}'  localhost:8080 helloworld.Greeter/SayHello
// {
//   "message": "hello world"
// }

type server struct {
	// Embed the unimplemented server
	helloworldpb.UnimplementedGreeterServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) SayHello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
	return &helloworldpb.HelloReply{Message: in.Name + " world"}, nil
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	helloworldpb.RegisterGreeterServer(s, &server{})

	reflection.Register(s)

	// Serve gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	log.Fatal(s.Serve(lis))
}
