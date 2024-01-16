package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mochi-yu/kns23-catch-up/app/usecase"
)

type HealthHandler interface {
	IndexGet(c *gin.Context)
	IndexPost(c *gin.Context)
}

type healthHandler struct {
	hu usecase.HealthUseCase
}

func NewHealthHandler(hu usecase.HealthUseCase) HealthHandler {
	return &healthHandler{hu: hu}
}

func (hh *healthHandler) IndexGet(c *gin.Context) {
}

func (hh *healthHandler) IndexPost(c *gin.Context) {
}
