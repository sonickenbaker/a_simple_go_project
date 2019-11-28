package main

import (
	"fmt"
	"os"
	"github.com/sonickenbaker/a_simple_go_project/hello"
)

func main() {
	if log_level := os.Getenv("DEBUG"); len(log_level) > 0 {
		fmt.Println(hello.SayHello(log_level))
	} else {
		fmt.Println(hello.SayHello("INFO"))
	}
}