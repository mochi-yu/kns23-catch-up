package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mochi-yu/kns23-catch-up/app/infrastructure"
)

type LoginAuth struct {
	fc *infrastructure.FirebaseClient
}

func NewLoginAuthMiddleware(fc *infrastructure.FirebaseClient) *LoginAuth {
	return &LoginAuth{fc: fc}
}

func (la *LoginAuth) Check() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Firebaseで認証トークンを取得
		auth_header := ctx.GetHeader("Authorization")
		idToken := strings.Replace(auth_header, "Bearer ", "", 1)
		token, err := la.fc.C.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			ctx.JSON(http.StatusForbidden, gin.H{"error": err})
			ctx.Abort()
			return
		}

		// コンテキストにuidを追加
		ctx.Set("uid", token.UID)
		fmt.Println(token.UID)
	}
}

func AdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
