package app

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mochi-yu/kns23-catch-up/app/handler"
	"github.com/mochi-yu/kns23-catch-up/app/infrastructure"
	"github.com/mochi-yu/kns23-catch-up/app/middleware"
	"github.com/mochi-yu/kns23-catch-up/app/repository"
	"github.com/mochi-yu/kns23-catch-up/app/usecase"
)

type Server struct {
	Engine     *gin.Engine
	repository *repository.Repository
	usecase    *usecase.UseCase
	handler    *handler.Handler
	middleware *middleware.Middleware
}

func NewServer() *Server {
	r := gin.Default()

	// ここからCorsの設定
	r.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{"*"},
		// アクセスを許可したいHTTPメソッド
		AllowMethods: []string{"*"},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{"*"},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))

	s := &Server{Engine: r}

	// DynamoDBの初期化
	dynamoDBClient := infrastructure.NewDynamoDBClient()

	// S3の初期化
	s3Client := infrastructure.NewS3Client()

	// firebaseの初期化
	firebaseCLient := infrastructure.NewFirebaseClient()

	// repositoryを初期化
	s.repository = &repository.Repository{
		S3Client: s3Client,
		User:     repository.NewUserRepository(*dynamoDBClient),
		Post:     repository.NewPostRepository(*dynamoDBClient),
	}

	// usecaseを初期化
	s.usecase = &usecase.UseCase{
		Health: usecase.NewHealthUseCase(s.repository),
		User:   usecase.NewUserUseCase(s.repository),
		Post:   usecase.NewPostUseCase(s.repository),
	}

	// ハンドラーを初期化
	s.handler = &handler.Handler{
		Health: handler.NewHealthHandler(s.usecase.Health),
		Post:   handler.NewPostHandler(s.usecase.Post),
		User:   handler.NewUserHandler(s.usecase.User),
	}

	// ミドルウェアを初期化
	s.middleware = &middleware.Middleware{
		LoginAuth: *middleware.NewLoginAuthMiddleware(firebaseCLient),
	}

	// ルーティングを定義
	s.setUpRouter()

	return s
}

func (s *Server) setUpRouter() {
	// ルーティングの定義
	// ログインを必要としないエンドポイント
	v1Group := s.Engine.Group("/v1")
	v1Group.GET("/health", s.handler.Health.IndexGet)
	v1Group.POST("/health", s.handler.Health.IndexPost)

	// ログインが必要なグループ
	needLoginGroup := v1Group.Group("/")
	needLoginGroup.Use(s.middleware.LoginAuth.Check())

	authGroup := needLoginGroup.Group("/auth")
	authGroup.POST("/", s.handler.User.PostAuth)
	authGroup.GET("/", s.handler.User.GetAuth)
	authGroup.POST("/temp", s.handler.User.PostAuthTemp)

	userGroup := needLoginGroup.Group("/users")
	userGroup.GET("/", s.handler.User.GetUsers)
	userGroup.GET("/{user_id}", s.handler.User.GetByUserID)
	userGroup.PUT("/{user_id}", s.handler.User.PutByUserID)
	userGroup.GET("/{user_id}/posts", s.handler.User.GetUserPosts)
	userGroup.GET("/{user_id}/reactions", s.handler.User.GetUserReactions)
	userGroup.GET("/{user_id}/comments", s.handler.User.GetUserComments)

	postGroup := needLoginGroup.Group("/posts")
	postGroup.GET("/", s.handler.Post.GetPosts)
	postGroup.POST("/", s.handler.Post.PostNewPost)
	postGroup.GET("/{post_id}", s.handler.Post.GetPostByPostID)
	postGroup.PUT("/{post_id}", s.handler.Post.PutPostByPostID)
	postGroup.DELETE("/{post_id}", s.handler.Post.DeletePostByPostID)
	postGroup.GET("/{post_id}/reactions", s.handler.Post.GetPostReactionsByPostID)
	postGroup.POST("/{post_id}/reactions", s.handler.Post.PostPostReactionsByPostID)
	postGroup.DELETE("/{post_id}/reactions", s.handler.Post.DeletePostReactionsByPostID)
	postGroup.GET("/{post_id}/comments", s.handler.Post.GetPostCommentsByPostID)
	postGroup.POST("/{post_id}/comments", s.handler.Post.PostPostCommentsByPostID)
	postGroup.DELETE("/{post_id}/comments", s.handler.Post.DeletePostCommentsByPostID)

	// admin権限が必要なグループ
	adminGroup := v1Group.Group("/admin")
	adminGroup.Use(middleware.AdminAuth())
}
