package main

import (
	"context"
	"fmt"
	"github.com/anz-bank/syslgen-examples/gen/todos"
	"net/http"
	"net/http/httptrace"
)

func withTrace(ctx context.Context) context.Context {
	trace := &httptrace.ClientTrace{
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %+v\n", dnsInfo)
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", connInfo.Conn.RemoteAddr())
		},
	}
	return httptrace.WithClientTrace(ctx, trace)
}

func main() {
	httpClient := http.Client{}
	client := todos.NewClient(&httpClient, "http://localhost:8080")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	result, err := client.GetTodosID(withTrace(ctx), map[string]string{}, 1)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	} else {
		fmt.Printf("Status: %d\nValue: \n%+v\n", result.HTTPResponse.StatusCode, result.Response)
	}
}
