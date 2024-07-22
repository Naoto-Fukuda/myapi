package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Naoto-Fukuda/myapi/apperrors"
	"github.com/Naoto-Fukuda/myapi/controllers/services"
	"github.com/Naoto-Fukuda/myapi/models"
)

type CommentController struct{
	service services.CommentServicer
}

func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

// POST /comment
func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		err = apperrors.RedBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to post data")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(comment)
}
