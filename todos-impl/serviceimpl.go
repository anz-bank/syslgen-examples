package todosimpl

import (
	"encoding/json"
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

func (s *ServiceImpl) GetTodosID(w http.ResponseWriter, id string) {
	for _, todo := range Todos {
		if strconv.Itoa(int(todo.ID)) == id {
			_ = json.NewEncoder(w).Encode(&todo)
			return
		}
	}
}

func (s *ServiceImpl) GetPosts(w http.ResponseWriter) {
	_ = json.NewEncoder(w).Encode(&PostList)
}

func (s *ServiceImpl) GetComments(w http.ResponseWriter, postID string) {
	var retValues todos.Posts
	for _, post := range PostList {
		if strconv.Itoa(int(post.ID)) == postID {
			retValues = append(retValues, post)
		}
	}
	_ = json.NewEncoder(w).Encode(&retValues)

}
func (s *ServiceImpl) PostComments(w http.ResponseWriter, newPost todos.Post) {
	PostList = append(PostList, newPost)
	_ = json.NewEncoder(w).Encode(&newPost)
	_ = json.NewEncoder(w).Encode(&PostList)
}
func (s *ServiceImpl) IsAuthorized(r *http.Request, authHeader string) bool {
	fmt.Println(authHeader)
	return true
}
func (s *ServiceImpl) SendErrorResponse(w http.ResponseWriter, statusCode int, message string, errObj error) {
	fmt.Println(statusCode)
	fmt.Println(message)
	fmt.Println(errObj != nil)
}
