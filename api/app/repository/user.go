package repository

import (
	"context"
	"fmt"

	"github.com/mochi-yu/kns23-catch-up/app/infrastructure"
	"github.com/mochi-yu/kns23-catch-up/app/model"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type UserRepository interface {
	InsertNewUser(u model.UserModel) error
}

type userRepository struct {
	db *infrastructure.DynamoDBClient
}

func NewUserRepository(dynamodb infrastructure.DynamoDBClient) UserRepository {
	return &userRepository{db: &dynamodb}
}

func (ur *userRepository) InsertNewUser(u model.UserModel) error {
	// DBのキーとマッピング
	av, err := attributevalue.MarshalMap(u)
	if err != nil {
		return fmt.Errorf("failed marshal db attribute: %v", err)
	}

	// インサート処理を実施
	if _, err = ur.db.C.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String("users"),
		Item:      av,
	}); err != nil {
		return fmt.Errorf("failed to put new user: %v", err)
	}

	return nil
}
