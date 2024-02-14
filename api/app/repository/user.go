package repository

import (
	"context"
	"fmt"

	"github.com/mochi-yu/kns23-catch-up/app/infrastructure"
	"github.com/mochi-yu/kns23-catch-up/app/model"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type UserRepository interface {
	InsertNewUser(u model.UserModel) error
	GetUserByUid(uid string) (
		*model.UserModel,
		error,
	)
	GetTempUserByUid(uid string) (
		*model.TempUserModel,
		error,
	)
	InsertTempUser(tu model.TempUserModel) error
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

func (ur *userRepository) GetUserByUid(uid string) (
	*model.UserModel,
	error,
) {
	res, err := ur.db.C.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("users"),
		Key: map[string]types.AttributeValue{
			"firebase_id": &types.AttributeValueMemberS{
				Value: uid,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	// ヒットするユーザがいなければ
	if res.Item == nil {
		return nil, fmt.Errorf("no user")
	}

	var tu model.UserModel
	if err := attributevalue.UnmarshalMap(res.Item, &tu); err != nil {
		return nil, err
	}

	return &tu, nil
}

func (ur *userRepository) GetTempUserByUid(uid string) (
	*model.TempUserModel,
	error,
) {
	res, err := ur.db.C.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("temp_users"),
		Key: map[string]types.AttributeValue{
			"firebase_id": &types.AttributeValueMemberS{
				Value: uid,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	// ヒットするユーザがいなければ
	if res.Item == nil {
		return nil, fmt.Errorf("no temp_user")
	}

	var tu model.TempUserModel
	if err := attributevalue.UnmarshalMap(res.Item, &tu); err != nil {
		return nil, err
	}

	return &tu, nil
}

func (ur *userRepository) InsertTempUser(tu model.TempUserModel) error {
	// DBのキーとマッピング
	av, err := attributevalue.MarshalMap(tu)
	if err != nil {
		return fmt.Errorf("failed marshal db attribute: %v", err)
	}

	// インサート処理を実施
	if _, err = ur.db.C.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String("temp_users"),
		Item:      av,
	}); err != nil {
		return fmt.Errorf("failed to put new temp_user: %v", err)
	}

	return nil
}
