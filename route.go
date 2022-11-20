package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	posts []Post
)

func init() {
	posts = []Post{{Id: 1, Title: "Title 1", Text: "Text 1"}}
}

func getPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	result, err := json.Marshal(posts)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error marshaling the Posts array"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
	fmt.Println("GET /posts hit!")

}

func addPost(resp http.ResponseWriter, req *http.Request) {
	var post Post
	resp.Header().Set("Content-type", "application/json")
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error marshaling the Posts array"}`))
		return
	}
	post.Id = len(posts) + 1
	posts = append(posts, post)
	resp.WriteHeader(http.StatusOK)
	result, err := json.Marshal(post)
	resp.Write(result)
}
