package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mochi-yu/kns23-catch-up/app/usecase"
)

type PostHandler interface {
	GetPosts(c *gin.Context)
	PostNewPost(c *gin.Context)
	GetPostByPostID(c *gin.Context)
	PutPostByPostID(c *gin.Context)
	DeletePostByPostID(c *gin.Context)
	GetPostReactionsByPostID(c *gin.Context)
	PostPostReactionsByPostID(c *gin.Context)
	DeletePostReactionsByPostID(c *gin.Context)
	GetPostCommentsByPostID(c *gin.Context)
	PostPostCommentsByPostID(c *gin.Context)
	DeletePostCommentsByPostID(c *gin.Context)
}

type postHandler struct {
	pu usecase.PostUseCase
}

func NewPostHandler(pu usecase.PostUseCase) PostHandler {
	return &postHandler{pu: pu}
}

func (ph *postHandler) GetPosts(c *gin.Context) {
	// TODO: implement
}

func (ph *postHandler) PostNewPost(c *gin.Context) {
	// TODO: implement
}

func (ph *postHandler) GetPostByPostID(c *gin.Context) {
	// TODO: implement
}

func (ph *postHandler) PutPostByPostID(c *gin.Context) {
	// TODO: implement
}

func (ph *postHandler) DeletePostByPostID(c *gin.Context) {
	// TODO: implement
}

func (ph *postHandler) GetPostReactionsByPostID(c *gin.Context) {
	// TODO: implement
}

func (ph *postHandler) PostPostReactionsByPostID(c *gin.Context) {
	// TODO: implement
}

func (ph *postHandler) DeletePostReactionsByPostID(c *gin.Context) {
	// TODO: implement
}

func (ph *postHandler) GetPostCommentsByPostID(c *gin.Context) {
	// TODO: implement
}

func (ph *postHandler) PostPostCommentsByPostID(c *gin.Context) {
	// TODO: implement
}

func (ph *postHandler) DeletePostCommentsByPostID(c *gin.Context) {
	// TODO: implement
}
