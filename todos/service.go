package main

import (
	"../lib"
	"context"
	"fmt"
	"net/http"
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
type Posts []*Post

// Todos ...
type Todos interface {
	GET_comments(ctx context.Context, postId int64) (Posts, error)
	GET_posts(ctx context.Context) (Posts, error)
	GET_todos_id(ctx context.Context, id int64) (Todo, error)
	POST_comments(ctx context.Context, newPost *Post) (Post, error)
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
func (s *TodosClient) GET_comments(ctx context.Context, postId int64) (Posts, error) {
	out := Posts{}
	err := restlib.DoGet(ctx, fmt.Sprintf("%s/comments?postId=%v", s.url, postId), &out)
	return out, err
}

// GET_posts ...
func (s *TodosClient) GET_posts(ctx context.Context) (Posts, error) {
	out := Posts{}
	err := restlib.DoGet(ctx, fmt.Sprintf("%s/posts", s.url), &out)
	return out, err
}

// GET_todos_id ...
func (s *TodosClient) GET_todos_id(ctx context.Context, id int64) (Todo, error) {
	out := Todo{}
	err := restlib.DoGet(ctx, fmt.Sprintf("%s/todos/%v", s.url, id), &out)
	return out, err
}

// POST_comments ...
func (s *TodosClient) POST_comments(ctx context.Context, newPost *Post) (Post, error) {
	out := Post{}
	err := restlib.DoPost(ctx, fmt.Sprintf("%s/comments", s.url), newPost, out)
	return out, err
}
