package main

import (
	"log"
	"net/http"
	"github.com/Naoto-Fukuda/myapi/handlers"
	"github.com/gorilla/mux"
)


func main() {
	r := mux.NewRouter()

	// http.HandleFunc("/hello", handlers.HelloHandler)
	// http.HandleFunc("/article", handlers.PostArticleHandler)
	// http.HandleFunc("/article/list", handlers.ArticleListHandler)
	// http.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	// http.HandleFunc("/article/nice", handlers.PostNiceHandler)
	// http.HandleFunc("/article/comment", handlers.PostCommentHandler)

	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
