package model

import (
	"time"

	"firebase.google.com/go/v4/auth"
)

// 通常ユーザ登録時のパラメータ
type RegisterPostRequest struct {
	UserID      string `json:"user_id"`
	DisplayName string `json:"display_name"`
	UserName    string `json:"user_name"`
	ClassID     string `json:"class_id"`
}

// ユーザ情報の概要のレスポンス
// ユーザ一覧のレスポンスなどに使用する
type UserResponseSummary struct {
	UserID      string `json:"user_id"`
	DisplayName string `json:"display_name"`
	UserName    string `json:"user_name"`
	ClassID     string `json:"class_id"`
	IsTeacher   bool   `json:"is_teacher"`
}

// ユーザ情報の詳細のレスポンス
// 本人へのレスポンスに使用する
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

type TempUserModel struct {
	FirebaseID  string `dynamodbav:"firebase_id"`
	MailAddress string `dynamodbav:"mail_address"`

	CreatedAt int64 `dynamodbav:"created_at"`
}

func RegisterPostParam2UserModel(u *auth.UserInfo, p RegisterPostRequest) UserModel {
	currentTime := time.Now().Unix()

	return UserModel{
		FirebaseID:  u.UID,
		UserID:      p.UserID,
		DisplayName: p.DisplayName,
		UserName:    p.UserName,
		ClassID:     p.ClassID,
		MailAddress: u.Email,

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

func TempUserModel2UserModel(tum *TempUserModel) *UserModel {
	return &UserModel{
		FirebaseID:  tum.FirebaseID,
		MailAddress: tum.MailAddress,
	}
}
