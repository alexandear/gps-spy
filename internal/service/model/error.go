package model

import "github.com/devchallenge/spy-api/internal/models"

func NewError(message string) *models.Error {
	return &models.Error{
		Code:    new(int64),
		Message: &message,
	}
}
