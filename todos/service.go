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
	GET_comments(ctx context.Context, headers map[string]string, postId int64) (*restlib.HttpResult, error)
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
func (s *TodosClient) GET_comments(ctx context.Context, headers map[string]string, postId int64) (*restlib.HttpResult, error) {
	required := []string{}
	responses := []interface{}{&Posts{}}
	u, err := url.Parse(fmt.Sprintf("%s/comments", s.url))
	if err != nil {
		return nil, err
	}

	q := u.Query()
	u.RawQuery = q.Encode()
	return restlib.DoHttpRequest(s.client, ctx, "GET", u.String(), nil, headers, required, responses)
}

// GET_posts ...
func (s *TodosClient) GET_posts(ctx context.Context, headers map[string]string) (*restlib.HttpResult, error) {
	required := []string{}
	responses := []interface{}{&Posts{}}
	u, err := url.Parse(fmt.Sprintf("%s/posts", s.url))
	if err != nil {
		return nil, err
	}

	q := u.Query()
	u.RawQuery = q.Encode()
	return restlib.DoHttpRequest(s.client, ctx, "GET", u.String(), nil, headers, required, responses)
}

// GET_todos_id ...
func (s *TodosClient) GET_todos_id(ctx context.Context, headers map[string]string, id int64) (*restlib.HttpResult, error) {
	required := []string{}
	responses := []interface{}{&Todo{}}
	u, err := url.Parse(fmt.Sprintf("%s/todos/%v", s.url, id))
	if err != nil {
		return nil, err
	}

	q := u.Query()
	u.RawQuery = q.Encode()
	return restlib.DoHttpRequest(s.client, ctx, "GET", u.String(), nil, headers, required, responses)
}

// POST_comments ...
func (s *TodosClient) POST_comments(ctx context.Context, headers map[string]string) (*restlib.HttpResult, error) {
	required := []string{}
	responses := []interface{}{&Post{}}
	u, err := url.Parse(fmt.Sprintf("%s/comments", s.url))
	if err != nil {
		return nil, err
	}

	q := u.Query()
	u.RawQuery = q.Encode()
	return restlib.DoHttpRequest(s.client, ctx, "POST", u.String(), nil, headers, required, responses)
}
