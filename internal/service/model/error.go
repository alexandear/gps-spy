package model

import "github.com/devchallenge/spy-api/internal/models"

func NewError(message string) *models.Error {
	return &models.Error{
		Message: &message,
	}
}
