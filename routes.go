package main

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/adykaaa/go-clean-rest/entity"
	"github.com/adykaaa/go-clean-rest/repository"
)

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

func getPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error getting the posts from FireBase"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)

}

func addPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error marshaling the Posts array"}`))
		return
	}
	post.ID = rand.Int63()
	repo.Save(&post)
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(post)
}
