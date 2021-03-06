package todos

//
//    THIS IS AUTOGENERATED BY sysl
//

import (
	"errors"
	"net/http"
)

// DefaultTodosImpl  ...
type DefaultTodosImpl struct {
}

// NewDefaultTodosImpl for Todos
func NewDefaultTodosImpl() *DefaultTodosImpl {
	return &DefaultTodosImpl{}
}

// ServiceInterface for Todos
type ServiceInterface interface {
	GetComments(PostID string) (int, map[string]string, *Posts)
	GetPosts() (map[string]string, *Posts, *ResourceNotFoundError, *ErrorResponse)
	GetTodosID(ID string) (map[string]string, *Todo, *ResourceNotFoundError, *ErrorResponse)
	PostComments(newPost Post) (int, map[string]string, *Post)
	IsAuthorized(r *http.Request, authHeader string) bool
	GetErrorResponse(statusCode int, message string, errObj error) interface{}
}

// nolint:gocritic
// GetComments ...
func (d *DefaultTodosImpl) GetComments(PostID string) (int, map[string]string, *Posts) {
	panic(errors.New("not implemented"))
}

// nolint:gocritic
// GetPosts ...
func (d *DefaultTodosImpl) GetPosts() (map[string]string, *Posts, *ResourceNotFoundError, *ErrorResponse) {
	panic(errors.New("not implemented"))
}

// nolint:gocritic
// GetTodosID ...
func (d *DefaultTodosImpl) GetTodosID(ID string) (map[string]string, *Todo, *ResourceNotFoundError, *ErrorResponse) {
	panic(errors.New("not implemented"))
}

// nolint:gocritic
// PostComments ...
func (d *DefaultTodosImpl) PostComments(newPost Post) (int, map[string]string, *Post) {
	panic(errors.New("not implemented"))
}

// nolint:gocritic
// IsAuthorized ...
func (d *DefaultTodosImpl) IsAuthorized(r *http.Request, authHeader string) bool {
	panic(errors.New("not implemented"))
}

// nolint:gocritic
// GetErrorResponse ...
func (d *DefaultTodosImpl) GetErrorResponse(statusCode int, message string, errObj error) interface{} {
	panic(errors.New("not implemented"))
}
