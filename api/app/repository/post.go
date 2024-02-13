package repository

import (
	"github.com/mochi-yu/kns23-catch-up/app/infrastructure"
)

type PostRepository interface {
}

type postRepository struct {
	db *infrastructure.DynamoDBClient
}

func NewPostRepository(dynamodb infrastructure.DynamoDBClient) PostRepository {
	return &postRepository{db: &dynamodb}
}
