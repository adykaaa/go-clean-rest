package controller

import (
	"encoding/json"
	"net/http"

	"github.com/adykaaa/go-clean-rest/entity"
	svcerror "github.com/adykaaa/go-clean-rest/errors"
	"github.com/adykaaa/go-clean-rest/service"
)

var (
	postService service.PostService = service.NewPostService()
)

type controller struct{}

type PostController interface {
	AddPost(resp http.ResponseWriter, req *http.Request)
	GetPosts(resp http.ResponseWriter, req *http.Request)
}

func NewPostController() PostController {
	return &controller{}
}

func (*controller) GetPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(svcerror.ServiceError{Message: "Error getting posts!"})
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)

}

func (*controller) AddPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	var post entity.Post
	decode_err := json.NewDecoder(req.Body).Decode(&post)
	if decode_err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(svcerror.ServiceError{Message: "Error unmarshaling post into JSON!"})
		return
	}

	validate_err := postService.Validate(&post)
	if validate_err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(svcerror.ServiceError{Message: validate_err.Error()})
		return
	}

	result, post_err := postService.Create(&post)
	if post_err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(svcerror.ServiceError{Message: "Error saving the post!"})
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(result)
}
