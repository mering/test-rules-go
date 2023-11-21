package main

import (
	"fmt"

	"src/msg_proto"
	"src/foo_proto"
	svc "src/svc_go"
)

func main() {
	foo := &foo_proto.Foo{}
	foo.Msg = &msg_proto.Msg{}
	foo.Msg.Payload = "world"
	fmt.Println("Hello", foo.Msg.Payload)

	svc.Run()
}
