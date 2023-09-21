package main

import (
	"fmt"

	"src/bar_go_proto"
	"src/foo_go_proto"
)

func main() {
	foo := &foo_go_proto.Foo{}
	foo.Bar = &bar_go_proto.Bar{}
	foo.Bar.X = 5
	fmt.Println("Hello world:", foo.Bar.X)
}
