package main

import (
	"fmt"

	"github.com/vishnu-g-k/to-do-list-application/controller"
)

func main() {
	fmt.Println("...Starting Application...")
	controller.TodoListHandler()
}
