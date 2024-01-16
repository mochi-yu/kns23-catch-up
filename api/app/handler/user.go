package handler

import (
	"github.com/gin-gonic/gin"
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
	// TODO: implement
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
