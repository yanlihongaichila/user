package main

import (
	"context"
	"fmt"
	"github.com/yanlihongaichila/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

//var addr = flag.String("addr", "localhost:8077", "the address to connect to")

func main() {

	creds, err := credentials.NewClientTLSFromFile("ca.pem", "x.test.example.com")

	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial("127.0.0.1:8077", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	userSvc := user.NewUserClient(conn)
	fmt.Println(userSvc)
	getUser, aaa := userSvc.GetUserByUsername(context.Background(), &user.GetUserByUsernameRequest{Username: "严礼鸿"})
	if aaa != nil {
		fmt.Println(aaa.Error())
		return
	}

	fmt.Println(aaa)
	fmt.Println(getUser)
	// Make a echo client and send an RPC.
	//rgc := ecpb.NewEchoClient(conn)
	//callUnaryEcho(rgc, "hello world")
}
