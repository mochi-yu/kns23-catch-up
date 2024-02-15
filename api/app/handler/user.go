package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mochi-yu/kns23-catch-up/app/model"
	"github.com/mochi-yu/kns23-catch-up/app/usecase"
	"github.com/mochi-yu/kns23-catch-up/app/util"
)

type UserHandler interface {
	PostAuth(c *gin.Context)
	GetAuth(c *gin.Context)
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

func (uh *userHandler) PostAuth(c *gin.Context) {
	// ユーザ情報を取得
	u, err := util.GetUserInfoFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		c.Abort()
		return
	}

	// パラメータを取得する
	requestParam := model.RegisterPostRequest{}
	if err := c.BindJSON(&requestParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		c.Abort()
	}

	// パラメータのバリデーション
	if requestParam.UserID == "" || requestParam.DisplayName == "" ||
		requestParam.UserName == "" || requestParam.ClassID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "パラメータの値が不足しています。"})
	}

	// メインの処理を実施
	userInfo, err := uh.uu.PostAuth(u, requestParam)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		c.Abort()
	}

	// レスポンスを返す
	response := model.UserModel2UserResponse(userInfo, false)
	c.JSON(http.StatusOK, response)
}

func (uh *userHandler) GetAuth(c *gin.Context) {
	// ユーザ情報を取得
	u, err := util.GetUserInfoFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		c.Abort()
		return
	}

	// メインの処理を実施
	userInfo, isTempUser, err := uh.uu.GetAuth(u)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		c.Abort()
		return
	}

	// レスポンスを返す
	response := model.UserModel2UserResponse(userInfo, isTempUser)
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
