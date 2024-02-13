package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mochi-yu/kns23-catch-up/app/model"
	"github.com/mochi-yu/kns23-catch-up/app/usecase"
)

type UserHandler interface {
	RegisterPost(c *gin.Context)
	GetUsers(c *gin.Context)
	GetByUserID(c *gin.Context)
	PutByUserID(c *gin.Context)
	GetUserPosts(c *gin.Context)
	GetUserReactions(c *gin.Context)
	GetUserComments(c *gin.Context)
}

type userHandler struct {
	uu usecase.UserUseCase
}

func NewUserHandler(uu usecase.UserUseCase) UserHandler {
	return &userHandler{uu: uu}
}

func (uh *userHandler) RegisterPost(c *gin.Context) {
	// パラメータを取得する
	requestParam := model.RegisterPostParam{}
	if err := c.BindJSON(&requestParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		c.Abort()
	}

	// パラメータのバリデーション
	if requestParam.UserID == "" || requestParam.DisplayName == "" ||
		requestParam.UserName == "" || requestParam.MailAddress == "" ||
		requestParam.ClassID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "パラメータの値が不足しています。"})
	}

	// メインの処理
	err := uh.uu.RegisterPost(requestParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		c.Abort()
	}

	// レスポンスを返す
	response := model.RegisterPostResponse{
		UserID:      requestParam.UserID,
		DisplayName: requestParam.DisplayName,
		ClassID:     requestParam.ClassID,
		MailAddress: requestParam.MailAddress,
	}
	c.JSON(http.StatusOK, response)
}

func (uh *userHandler) GetUsers(c *gin.Context) {
	// TODO: implement
}

func (uh *userHandler) GetByUserID(c *gin.Context) {
	// TODO: implement
}

func (uh *userHandler) PutByUserID(c *gin.Context) {
	// TODO: implement
}

func (uh *userHandler) GetUserPosts(c *gin.Context) {
	// TODO: implement
}

func (uh *userHandler) GetUserReactions(c *gin.Context) {
	// TODO: implement
}

func (uh *userHandler) GetUserComments(c *gin.Context) {
	// TODO: implement
}
