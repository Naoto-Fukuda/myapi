package services

import (
	"github.com/Naoto-Fukuda/myapi/models"
	"github.com/Naoto-Fukuda/myapi/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return newComment, nil
}
