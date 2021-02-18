package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type PostData struct {
	UserId int `json:"userId"`
	Id int `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
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

	go func () {
		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gogolang")
		if err != nil {
			panic(err)
		}
		defer db.Close()
	
		for _, i := range dataFromPost { 
			q := "INSERT INTO `posts`(`userId`, `id`, `title`, `body`) VALUES (?,?,?,?)"
			insert, err := db.Prepare(q)
			defer insert.Close()
			if err != nil {
				panic(err)
			}

			_, err = insert.Exec(i.UserId, i.Id, i.Title, i.Body)
			if err != nil {
				panic(err)
			}

		}
		
	}()
		
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

	go func () {
		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gogolang")
		if err != nil {
			panic(err)
		}
		defer db.Close()
	
		for _, i := range dataFromComment { 
			q := "INSERT INTO `comments`(`postId`, `id`, `name`, `email`, `body`) VALUES (?,?,?,?,?)"
			insert, err := db.Prepare(q)
			defer insert.Close()
			if err != nil {
				panic(err)
			}

			_, err = insert.Exec(i.PostId, i.Id, i.Name, i.Email, i.Body)
			if err != nil {
				panic(err)
			}

		}
		
	}()
	
}

func main() {

	go getPostConn(7)

	var input string
	fmt.Scanln(&input)
}