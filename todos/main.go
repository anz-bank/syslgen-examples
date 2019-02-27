package main

import (
	"context"
	"fmt"
	"net/http"
)

func main() {
	httpClient := http.Client{}
	client := MakeTodosClient(&httpClient, "http://jsonplaceholder.typicode.com")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	res, err := client.GET_todos_id(ctx, 1)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	} else {
		fmt.Printf("Value: \n%+v\n", res)
	}
}
