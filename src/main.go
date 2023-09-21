package main

import (
	"fmt"

	"src/bar_proto"
	"src/foo_proto"
)

func main() {
	foo := &foo_proto.Foo{}
	foo.Bar = &bar_proto.Bar{}
	foo.Bar.X = 5
	fmt.Println("Hello world:", foo.Bar.X)
}
