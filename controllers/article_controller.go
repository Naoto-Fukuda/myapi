package controllers

import (
	"encoding/json"
	// "errors"
	// "fmt"
	// "io"
	// "log"
	"net/http"
	"strconv"

	"github.com/Naoto-Fukuda/myapi/models"
	"github.com/Naoto-Fukuda/myapi/controllers/services"
	"github.com/gorilla/mux"
)

type ArticleController struct {
	service services.ArticleServicer
}

// コンストラクタ関数
func NewArticleController(s services.ArticleServicer) *ArticleController {
	return &ArticleController{service: s}
}

func NewMyAppController(s services.ArticleServicer) *ArticleController {
	return NewArticleController(s)
}

// POST /article
func (c *ArticleController)PostArticleHandler(w http.ResponseWriter, req *http.Request){	
	var reqArticle models.Article

	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	article, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	// ストリームからメモリ変換せずにreq.Bodyを元に構造体を作成
	json.NewEncoder(w).Encode(article)
}

// GET /article/list
func (c *ArticleController) ArticleListHandler(w http.ResponseWriter, req *http.Request){
	queryMap := req.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		} 
	} else {
			page = 1
	}

	articleList, err := c.service.GetArticleListService(page)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(articleList)
}

// GET /article/{id}
func (c *ArticleController) ArticleDetailHandler(w http.ResponseWriter, req *http.Request){
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
	}

	article, err := c.service.GetArticleService(articleID)
	if err != nil {
		http.Error(w, "fatal internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

// POST /article/nice 
func (c *ArticleController) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	article, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		http.Error(w, "fatal internal exec\n", http.StatusInternalServerError)
		return
	}
	
	json.NewEncoder(w).Encode(article)
}
