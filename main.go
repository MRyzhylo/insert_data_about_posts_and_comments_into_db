package main

import (
	"fmt"
	"net/http"
	"log"
	"strconv"
	"encoding/json"
	"io/ioutil"
	// "bytes"
)

type PostData struct {
	UserId int `json:"userId"`
	Id int `json:"id"`
	Title string `json:"title"`
	Body string `json"body"`
}

type CommentData struct {
	PostId int `json:"postId"`
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Body string `json:"body"`
}

func getPostConn(userId int) {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts?userId=" + strconv.Itoa(userId))
	if err != nil {
		log.Fatal(err)
	}

	post, _ := ioutil.ReadAll(res.Body)

	var dataFromPost []PostData
	err = json.Unmarshal(post, &dataFromPost)
		
	for _, x := range dataFromPost {
		go getCommentFromPost(x.Id)
	}
}

func getCommentFromPost(postId int) {
	res, err := http.Get("https://jsonplaceholder.typicode.com/comments?postId=" + strconv.Itoa(postId))
	if err != nil {
		log.Fatal(err)
	}

	comment, _ := ioutil.ReadAll(res.Body)

	var dataFromComment []CommentData
	err = json.Unmarshal(comment, &dataFromComment)

	fmt.Printf("%v", dataFromComment)
}

func main() {

	go getPostConn(7)

	var input string
	fmt.Scanln(&input)
}