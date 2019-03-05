package main

import (
	"../lib"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// Post ...
type Post struct {
	Body   string `json:"body"`
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Userid int64  `json:"userid"`
}

// Todo ...
type Todo struct {
	Completed bool   `json:"completed"`
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Userid    int64  `json:"userid"`
}

// Posts ...
type Posts []Post

// Todos ...
type Todos interface {
	GET_comments(ctx context.Context, headers map[string]string, postId string) (*restlib.HttpResult, error)
	GET_posts(ctx context.Context, headers map[string]string) (*restlib.HttpResult, error)
	GET_todos_id(ctx context.Context, headers map[string]string, id int64) (*restlib.HttpResult, error)
	POST_comments(ctx context.Context, headers map[string]string) (*restlib.HttpResult, error)
}

// TodosClient ...
type TodosClient struct {
	client *http.Client
	url    string
}

// MakeTodosClient ...
func MakeTodosClient(client *http.Client, url string) *TodosClient {
	return &TodosClient{client, url}
}

// GET_comments ...
func (s *TodosClient) GET_comments(ctx context.Context, headers map[string]string, postId string) (*restlib.HttpResult, error) {
	required := []string{}
	responses := []interface{}{&Posts{}}
	return restlib.DoHttpRequest(s.client, ctx, "GET", fmt.Sprintf("%s/comments?postId=%v", s.url, url.QueryEscape(postId)), nil, headers, required, responses)
}

// GET_posts ...
func (s *TodosClient) GET_posts(ctx context.Context, headers map[string]string) (*restlib.HttpResult, error) {
	required := []string{}
	responses := []interface{}{&Posts{}}
	return restlib.DoHttpRequest(s.client, ctx, "GET", fmt.Sprintf("%s/posts", s.url), nil, headers, required, responses)
}

// GET_todos_id ...
func (s *TodosClient) GET_todos_id(ctx context.Context, headers map[string]string, id int64) (*restlib.HttpResult, error) {
	required := []string{}
	responses := []interface{}{&Todo{}}
	return restlib.DoHttpRequest(s.client, ctx, "GET", fmt.Sprintf("%s/todos/%v", s.url, id), nil, headers, required, responses)
}

// POST_comments ...
func (s *TodosClient) POST_comments(ctx context.Context, headers map[string]string) (*restlib.HttpResult, error) {
	required := []string{}
	responses := []interface{}{&Post{}}
	return restlib.DoHttpRequest(s.client, ctx, "POST", fmt.Sprintf("%s/comments", s.url), nil, headers, required, responses)
}
