package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptrace"

	"../todos"
)

func withTrace(context context.Context) context.Context {
	trace := &httptrace.ClientTrace{
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %+v\n", dnsInfo)
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", connInfo.Conn.RemoteAddr())
		},
	}
	return httptrace.WithClientTrace(context, trace)
}

func main() {
	httpClient := http.Client{}
	client := todos.MakeTodosClient(&httpClient, "http://jsonplaceholder.typicode.com")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	result, err := client.GETTodosID(withTrace(ctx), map[string]string{}, 1)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	} else {
		fmt.Printf("Value: \n%+v\n", result.Response)
	}
}
