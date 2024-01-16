package middleware

import (
	"github.com/gin-gonic/gin"
)

func LoginAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// auth_header := ctx.GetHeader("Authorization")
		// idToken := strings.Replace(auth_header, "Bearer ", "", 1)
		// token := r.firebaseApp.VerifyIDToken(idToken)
		// uid := token.UID
		// user, err := gateways.NewUserRepository(r.db).FindByUid(uid)
		// if err != nil {
		// 	fmt.Print("不正なリクエスト")
		// 	ctx.JSON(401, gin.H{"message": "認証に失敗しました"})
		// } else {
		// 	ctx.Next()
		// }
	}
}

func AdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
