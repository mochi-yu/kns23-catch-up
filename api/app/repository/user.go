package repository

import "github.com/mochi-yu/kns23-catch-up/app/infrastructure"

type UserRepository interface {
}

type userRepository struct {
	db *infrastructure.DynamoDBClient
}

func NewUserRepository(dynamodb *infrastructure.DynamoDBClient) UserRepository {
	return &userRepository{db: dynamodb}
}
