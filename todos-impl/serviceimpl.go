package todosimpl

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/anz-bank/syslgen-examples/todos"
)

// ServiceImpl for Todos API
type ServiceImpl struct {
}

// NewServiceImpl for Todos
func NewServiceImpl() *ServiceImpl {
	return &ServiceImpl{}
}

func (s *ServiceImpl) GetComments(postID string) (int, map[string]string, *todos.Posts) {
	var retValues todos.Posts
	for _, post := range PostList {
		if strconv.Itoa(int(post.ID)) == postID {
			retValues = append(retValues, post)
		}
	}
	headers := map[string]string{}
	return http.StatusOK, headers, &retValues
}

func (s *ServiceImpl) GetPosts() (int, map[string]string, *todos.Posts) {
	headers := map[string]string{}
	return http.StatusOK, headers, &PostList
}

func (s *ServiceImpl) GetTodosID(id string) (int, map[string]string, *todos.Todo) {
	headers := map[string]string{}
	for _, todo := range Todos {
		if strconv.Itoa(int(todo.ID)) == id {
			return http.StatusOK, headers, &todo
		}
	}

	return http.StatusNotFound, headers, nil
}

func (s *ServiceImpl) PostComments(newPost todos.Post) (int, map[string]string, *todos.Post) {
	PostList = append(PostList, newPost)
	headers := map[string]string{}
	return http.StatusOK, headers, &newPost
}

func (s *ServiceImpl) IsAuthorized(r *http.Request, authHeader string) bool {
	fmt.Println(authHeader)
	return true
}

func (s *ServiceImpl) GetErrorResponse(statusCode int, message string, errObj error) interface{} {
	fmt.Println("GetErrorResponse statusCode: ", statusCode)
	fmt.Println("GetErrorResponse msg: ", message)
	fmt.Println("GetErrorResponse errObj is null?: ", errObj != nil)
	return nil
}
