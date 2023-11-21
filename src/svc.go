package svc_go

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	msg "src/msg_proto"
	svc "src/svc_go_grpc"
)

func Run() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("127.0.0.1:10101", opts...)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	client := svc.NewGreeterClient(conn)

	feature, err := client.SayHello(context.Background(), &msg.Msg{Payload: "Hello world!"})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	log.Println(feature.Payload)
}
