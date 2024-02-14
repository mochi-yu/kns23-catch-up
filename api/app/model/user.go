package model

import "time"

type RegisterPostParam struct {
	UserID      string `json:"userID"`
	DisplayName string `json:"displayName"`
	UserName    string `json:"userName"`
	ClassID     string `json:"ClassID"`
	MailAddress string `json:"mailAddress"`
}

type RegisterPostResponse struct {
type UserResponse struct {
	UserID      string `json:"user_id"`
	DisplayName string `json:"display_name"`
	UserName    string `json:"user_name"`
	ClassID     string `json:"class_id"`
	MailAddress string `json:"mail_address"`

	IsTempUser bool `json:"is_temp_user"`
	IsAdmin    bool `json:"is_admin"`
	IsTeacher  bool `json:"is_teacher"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

type UserModel struct {
	FirebaseID  string `dynamodbav:"firebase_id"`
	UserID      string `dynamodbav:"user_id"`
	DisplayName string `dynamodbav:"display_name"`
	UserName    string `dynamodbav:"user_name"`
	ClassID     string `dynamodbav:"class_id"`
	MailAddress string `dynamodbav:"mail_address"`

	IsAdmin   bool `dynamodbav:"is_admin"`
	IsTeacher bool `dynamodbav:"is_teacher"`

	CreatedAt int64 `dynamodbav:"created_at"`
	UpdatedAt int64 `dynamodbav:"updated_at"`
}

func RegisterParam2UserModel(param RegisterPostParam) *UserModel {
	currentTime := time.Now().Unix()

	return &UserModel{
		UserID:      param.UserID,
		DisplayName: param.DisplayName,
		UserName:    param.UserName,
		ClassID:     param.ClassID,
		MailAddress: param.MailAddress,
		FirebaseID:  "test",

		IsAdmin:   false,
		IsTeacher: false,

		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}
}

func UserModel2UserResponse(um *UserModel, isTempUser bool) UserResponse {
	return UserResponse{
		UserID:      um.UserID,
		DisplayName: um.DisplayName,
		UserName:    um.UserName,
		ClassID:     um.ClassID,
		MailAddress: um.MailAddress,

		IsTempUser: isTempUser,
		IsAdmin:    um.IsAdmin,
		IsTeacher:  um.IsTeacher,

		CreatedAt: um.CreatedAt,
		UpdatedAt: um.UpdatedAt,
	}
}
