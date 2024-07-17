package handlers

import (
	"encoding/json"
	// "errors"
	// "fmt"
	// "io"
	"log"
	"net/http"
	"strconv"

	"github.com/Naoto-Fukuda/myapi/models"
	"github.com/gorilla/mux"
)

// POST /article
func PostArticleHandler(w http.ResponseWriter, req *http.Request){

	// Header.GETの返り値はstringなのでAtoiで整数に変換
	// length, err := strconv.Atoi(req.Header.Get("Content-Length"))
	// if err != nil {
	// 	http.Error(w, "cannot get content length\n", http.StatusBadRequest)
	// 	return
	// }

	// reqBodyBuffer := make([]byte, length)

	// if _, err := req.Body.Read(reqBodyBuffer); !errors.Is(err, io.EOF) {
	// 	http.Error(w, "fail to get request body\n", http.StatusBadRequest)
	// 	return
	// }

	// defer req.Body.Close()
	
	var reqArticle models.Article

	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	article := reqArticle

	// if err := json.Unmarshal(reqBodyBuffer, &reqArticle); err != nil {
	// 	http.Error(w, "fail to decode json≠\n", http.StatusBadRequest)
	// 	return
	// }
	// if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
	// 	http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	// }

	// jsonData, err := json.Marshal(reqArticle)
	// if err != nil {
	// 	http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
	// 	return
	// }

	// w.Write(jsonData)

	// ストリームからメモリ変換せずにreq.Bodyを元に構造体を作成
	json.NewEncoder(w).Encode(article)
}

// GET /article/list
func ArticleListHandler(w http.ResponseWriter, req *http.Request){
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

	// 暫定でこれを追加することで
	// 「変数pageが使われていない」というコンパイルエラーを回避
	log.Println(page)
	articleList := []models.Article{models.Article1, models.Article2}
	json.NewEncoder(w).Encode(articleList)
}

// GET /article/{id}
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request){
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
	}
	// 暫定でこれを追加することで
	// 「変数articleIDが使われていない」というコンパイルエラーを回避
	log.Println(articleID)
	article := models.Article1
	json.NewEncoder(w).Encode(article)
}

// POST /article/nice 
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	article := reqArticle
	json.NewEncoder(w).Encode(article)
}

// POST /comment
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	comment := reqComment
	json.NewEncoder(w).Encode(comment)
}
