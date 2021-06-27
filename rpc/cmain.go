package rpc

import (
	"context"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "golangpratice/rpc/proto"
)

const (
	address     = ":5001"
	defaultName = "world"
)

var (
	clientSends = []pb.SayWhat{{What: "aaa"},
		{What: "bbb"},
		{What: "ccc"},
		{What: "ddd"},
		{What: "eee"}}
)

func say3(client pb.GreeterClient) {
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	log.Print("sending......")
	r, err := client.Say3(ctx, &pb.HelloRequest{Name: "hello"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}

// send and receive on stream
func say4(client pb.GreeterClient) {

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()
	stream, err := client.Say4(ctx)
	if err != nil {
		log.Println("SayStream error", err)
		return
	}

	wait := make(chan struct{})
	go func() {
		for {
			reply, err := stream.Recv()
			if err == io.EOF {
				log.Println("Recv EOF", err)
				close(wait)
				return
			}
			if err != nil {
				log.Println("Recv err", err)
				close(wait)
				return
			}

			log.Println("Recv", reply.GetWhat())
		}
	}()

	for _, s := range clientSends {
		if err := stream.Send(&s); err != nil {
			log.Println("Error send", err)
			return
		}
		log.Println("sent", s.GetWhat())
	}
	//stream.CloseSend()

	<-wait
	log.Println("exit wait")
}

// receive on stream
func say1(client pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.Say1(ctx, &clientSends[0])
	if err != nil {
		log.Fatal("error", err)
	}
	defer stream.CloseSend()

	for {
		reply, err := stream.Recv()
		if err == io.EOF {
			log.Println("EOF")
			return
		}
		if err != nil {
			log.Println("error", err)
			return
		}

		log.Println("received", reply.What)
	}

}

// send with stream
func say2(client pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.Say2(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range clientSends {
		if err := stream.Send(&s); err != nil {
			log.Println("error", err)
			return
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("received", reply)

}

func RPCClientMain(op string, tls bool) {
	log.Print("RPCClientMain entry......")
	var opts []grpc.DialOption

	if tls {
		cred, err := clientTLSLoad("./cert/client.crt", "./cert/client.key")
		if err != nil {
			log.Fatal("err")
		}
		log.Println("TLS enabled.")

		opts = append(opts, grpc.WithTransportCredentials(cred))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	opts = append(opts, grpc.WithBlock())

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	switch op {
	case "0":
		say3(c)
	case "1":
		say4(c)
	case "2":
		say1(c)
	case "3":
		say2(c)
	default:
		log.Fatal("Not supported!!!")
	}

}
