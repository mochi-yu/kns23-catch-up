package middleware

import (
	"context"
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

		// ユーザ情報を取得する
		u, err := la.fc.C.GetUser(context.Background(), token.UID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			ctx.Abort()
			return
		}

		// コンテキストにuser_infoを追加
		ctx.Set("user_info", u.UserInfo)

		ctx.Next()
	}
}

func AdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
