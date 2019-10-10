package main

import (
	"fmt"
	"github.com/apache/qpid-proton/go/pkg/amqp"
)

func main() {

	amqpe := amqp.Error{
		Name:        "myError",
		Description: "myDesc",
	}

	fmt.Println("My amqp Error is " + amqpe.Error())
}
