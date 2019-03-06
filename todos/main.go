package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/httptrace"
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
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := http.Client{
		Transport: tr,
	}

	client := MakeTodosClient(&httpClient, "http://jsonplaceholder.typicode.com")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	result, err := client.GET_todos_id(withTrace(ctx), map[string]string{}, 1)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	} else {
		fmt.Printf("Value: \n%+v\n", result.Response)
	}
}
