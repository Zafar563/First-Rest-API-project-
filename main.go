package main

import (
	"fmt"
	"todolist/httpserver"
	"todolist/todo"
)

func main() {
	todoList := todo.NewList()
	httpHandlers := httpserver.NewHTTPhandler(todoList)
	httpServer := httpserver.NewHTTPServer(httpHandlers)
	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start http server")
	}
}
