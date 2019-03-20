package todos

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/anz-bank/syslgen-examples/lib"
	"github.com/rickb777/date"
)

// Reference imports to suppress unused errors
var _ = time.Parse

// Reference imports to suppress unused errors
var _ = date.Parse

// Post ...
type Post struct {
	Body   string `json:"body"`
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	UserID int64  `json:"userId"`
}

// Todo ...
type Todo struct {
	Completed bool   `json:"completed"`
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	UserID    int64  `json:"userId"`
}

// Posts ...
type Posts []Post

// Todos ...
type Todos interface {
	GETComments(ctx context.Context, headers map[string]string, postId int64) (*restlib.HTTPResult, error)
	GETPosts(ctx context.Context, headers map[string]string) (*restlib.HTTPResult, error)
	GETTodosID(ctx context.Context, headers map[string]string, id int64) (*restlib.HTTPResult, error)
	POSTComments(ctx context.Context, headers map[string]string) (*restlib.HTTPResult, error)
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

// GETComments ...
func (s *TodosClient) GETComments(ctx context.Context, headers map[string]string, postId int64) (*restlib.HTTPResult, error) {
	required := []string{}
	responses := []interface{}{&Posts{}}
	u, err := url.Parse(fmt.Sprintf("%s/comments", s.url))
	if err != nil {
		return nil, err
	}

	q := u.Query()
	u.RawQuery = q.Encode()
	return restlib.DoHTTPRequest(ctx, s.client, "GET", u.String(), nil, headers, required, responses)
}

// GETPosts ...
func (s *TodosClient) GETPosts(ctx context.Context, headers map[string]string) (*restlib.HTTPResult, error) {
	required := []string{}
	responses := []interface{}{&Posts{}}
	u, err := url.Parse(fmt.Sprintf("%s/posts", s.url))
	if err != nil {
		return nil, err
	}

	q := u.Query()
	u.RawQuery = q.Encode()
	return restlib.DoHTTPRequest(ctx, s.client, "GET", u.String(), nil, headers, required, responses)
}

// GETTodosID ...
func (s *TodosClient) GETTodosID(ctx context.Context, headers map[string]string, id int64) (*restlib.HTTPResult, error) {
	required := []string{}
	responses := []interface{}{&Todo{}}
	u, err := url.Parse(fmt.Sprintf("%s/todos/%v", s.url, id))
	if err != nil {
		return nil, err
	}

	q := u.Query()
	u.RawQuery = q.Encode()
	return restlib.DoHTTPRequest(ctx, s.client, "GET", u.String(), nil, headers, required, responses)
}

// POSTComments ...
func (s *TodosClient) POSTComments(ctx context.Context, headers map[string]string) (*restlib.HTTPResult, error) {
	required := []string{}
	responses := []interface{}{&Post{}}
	u, err := url.Parse(fmt.Sprintf("%s/comments", s.url))
	if err != nil {
		return nil, err
	}

	q := u.Query()
	u.RawQuery = q.Encode()
	return restlib.DoHTTPRequest(ctx, s.client, "POST", u.String(), nil, headers, required, responses)
}
