package util

import (
	"fmt"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

func GetUserInfoFromContext(c *gin.Context) (*auth.UserInfo, error) {
	v, ok := c.Get("user_info")
	if !ok {
		return nil, fmt.Errorf("context not exist")
	}

	u, ok := v.(*auth.UserInfo)
	if !ok {
		return nil, fmt.Errorf("context not work well")
	}

	return u, nil
}
