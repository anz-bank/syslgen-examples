package todos

//
//    THIS IS AUTOGENERATED BY sysl
//

import (
	"encoding/json"
	"net/http"

	"github.com/anz-bank/syslgen-examples/pkg/restlib"
)

// Handler interface for Todos
type Handler interface {
	GetCommentsHandler(w http.ResponseWriter, r *http.Request)
	GetPostsHandler(w http.ResponseWriter, r *http.Request)
	GetTodosIDHandler(w http.ResponseWriter, r *http.Request)
	PostCommentsHandler(w http.ResponseWriter, r *http.Request)
}

// ServiceHandler for Todos API
type ServiceHandler struct {
	serviceInterface ServiceInterface
}

// NewServiceHandler for Todos
func NewServiceHandler(serviceInterface ServiceInterface) *ServiceHandler {
	return &ServiceHandler{serviceInterface}
}

// GetCommentsHandler ...
func (s *ServiceHandler) GetCommentsHandler(w http.ResponseWriter, r *http.Request) {
	PostID := restlib.GetQueryParam(r, "postId")
	Status, headerMap, Posts := s.serviceInterface.GetComments(PostID)
	restlib.SetHeaders(w, headerMap)
	restlib.SendHTTPResponse(w, Status, Posts)
}

// GetPostsHandler ...
func (s *ServiceHandler) GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	headerMap, Posts, ResourceNotFoundError, ErrorResponse := s.serviceInterface.GetPosts()
	var httpStatus int

	restlib.SetHeaders(w, headerMap)
	if Posts != nil {
		httpStatus = 200
		restlib.SendHTTPResponse(w, httpStatus, Posts)
		return
	}

	if ResourceNotFoundError != nil {
		httpStatus = 404
		errResp := s.serviceInterface.GetErrorResponse(httpStatus, "Internal server error", nil)
		restlib.SendHTTPResponse(w, httpStatus, errResp)
		return
	}

	if ErrorResponse != nil {
		httpStatus = 500
		errResp := s.serviceInterface.GetErrorResponse(httpStatus, "Internal server error", nil)
		restlib.SendHTTPResponse(w, httpStatus, errResp)
		return
	}

}

// GetTodosIDHandler ...
func (s *ServiceHandler) GetTodosIDHandler(w http.ResponseWriter, r *http.Request) {
	ID := restlib.GetURLParam(r, "id")
	headerMap, Todo, ResourceNotFoundError, ErrorResponse := s.serviceInterface.GetTodosID(ID)
	var httpStatus int

	restlib.SetHeaders(w, headerMap)
	if Todo != nil {
		httpStatus = 200
		restlib.SendHTTPResponse(w, httpStatus, Todo)
		return
	}

	if ResourceNotFoundError != nil {
		httpStatus = 404
		errResp := s.serviceInterface.GetErrorResponse(httpStatus, "Internal server error", nil)
		restlib.SendHTTPResponse(w, httpStatus, errResp)
		return
	}

	if ErrorResponse != nil {
		httpStatus = 500
		errResp := s.serviceInterface.GetErrorResponse(httpStatus, "Internal server error", nil)
		restlib.SendHTTPResponse(w, httpStatus, errResp)
		return
	}

}

// PostCommentsHandler ...
func (s *ServiceHandler) PostCommentsHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var newPost Post

	decodeErr := decoder.Decode(&newPost)
	if decodeErr != nil {
		errResp := s.serviceInterface.GetErrorResponse(http.StatusBadRequest, "Error reading request body", decodeErr)
		restlib.SendHTTPResponse(w, http.StatusBadRequest, errResp)
		return
	}

	Status, headerMap, Post := s.serviceInterface.PostComments(newPost)
	restlib.SetHeaders(w, headerMap)
	restlib.SendHTTPResponse(w, Status, Post)
}