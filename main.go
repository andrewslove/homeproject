package main

import (
	"fmt"
	"myproject/data"
	"myproject/valid"

	"github.com/andrewslove/homeproject/handlers"
)

func main() {
	fmt.Println("Hello World")
	handlers.Sayhello()
	fmt.Println(data.UserName())

	fmt.Println(valid.ValidateName(data.UserName()))
}
