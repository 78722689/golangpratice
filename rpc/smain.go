package rpc

import (
	"context"
	"io"
	"log"
	"net"

	pb "golangpratice/rpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	PORT = ":5001"
)

var (
	sends = []pb.ReplyWhat{{What: "111"},
		{What: "222"},
		{What: "333"},
		{What: "444"},
		{What: "555"}}
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) Say3(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println("request: ", in.Name)

	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) Say4(stream pb.Greeter_Say4Server) error {

	defer log.Println("exit SayStream...")

	//go func() error {
	defer log.Println("exit routine...")
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			log.Println("recv EOF.")
			return nil
		}
		if err != nil {
			log.Println("occurred error", err)
			return err
		}

		log.Println("Recieved", msg.GetWhat())
		for _, s := range sends {
			if err := stream.Send(&s); err != nil {
				log.Println("Send error", err)
				return err
			}
		}

		return nil
	}
	//}()

	//time.Sleep(10 * time.Second)
	return nil
}

func (s *server) Say1(in *pb.SayWhat, stream pb.Greeter_Say1Server) error {
	log.Println("Received", in.What)

	for _, s := range sends {
		if err := stream.Send(&s); err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}

func (s *server) Say2(stream pb.Greeter_Say2Server) error {

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			log.Println("EOF")
			return stream.SendAndClose(&pb.ReplyWhat{What: "xyz"})
		}
		if err != nil {
			log.Println("error", err)
			return err
		}

		log.Println("Received", message)
	}
	return nil
}

func RPCServerMain(tls bool) {
	var opts []grpc.ServerOption

	if tls {
		log.Println("TLS enabled.")
		cred, err := serverTLSLoad("./cert/localhost.crt", "./cert/localhost.key")
		if err != nil {
			log.Fatal(err)
		}
		opts = append(opts, grpc.Creds(cred))
	}

	listen, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	rpc := grpc.NewServer(opts...)
	pb.RegisterGreeterServer(rpc, &server{})
	reflection.Register(rpc)
	log.Println("RPC service is started")
	rpc.Serve(listen)
}
