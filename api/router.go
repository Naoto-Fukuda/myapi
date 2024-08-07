package api

import (
	"database/sql"
	"net/http"

	"github.com/Naoto-Fukuda/myapi/controllers"
	"github.com/Naoto-Fukuda/myapi/services"
	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	ser		:= services.NewMyAppService(db)
	aCon 	:= controllers.NewArticleController(ser)
	cCon 	:= controllers.NewCommentController(ser)

	r.HandleFunc("/article", aCon.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", aCon.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", aCon.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", cCon.PostCommentHandler).Methods(http.MethodPost)

	return r
}
