package main

import (
	"log"
	"net/http"

	"github.com/anz-bank/syslgen-examples/todos"
)

func main() {
	serviceImpl := todos.NewServiceImpl()
	svcHandler := todos.NewServiceHandler(serviceImpl)
	serviceRouter := todos.NewServiceRouter(svcHandler)
	router := serviceRouter.Route()

	log.Fatal(http.ListenAndServe(":8080", router))
}
