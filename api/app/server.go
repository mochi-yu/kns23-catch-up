package app

import (
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
}

func NewServer() *Server {
	r := gin.Default()
	s := &Server{Engine: r}

	// DynamoDBの初期化
	dynamoDBClient := infrastructure.NewDynamoDBClient()

	// S3の初期化
	s3Client := infrastructure.NewS3Client()

	// repositoryを初期化
	s.repository = &repository.Repository{
		S3Client: s3Client,
		User:     repository.NewUserRepository(dynamoDBClient),
		Post:     repository.NewPostRepository(dynamoDBClient),
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

	// ルーティングを定義
	s.setUpRouter(dynamoDBClient, s3Client)

	return s
}

func (s *Server) setUpRouter(
	dynamoDBClient infrastructure.DynamoDBClient,
	s3Client infrastructure.S3Client,
) {
	// ルーティングの定義
	// ログインを必要としないエンドポイント
	v1Group := s.Engine.Group("/v1")
	v1Group.GET("/health", s.handler.Health.IndexGet)
	v1Group.POST("/health", s.handler.Health.IndexPost)
	v1Group.POST("/register", s.handler.User.RegisterPost)

	// ログインが必要なグループ
	authGroup := v1Group.Group("/")
	authGroup.Use(middleware.LoginAuth())

	userGroup := authGroup.Group("/users")
	userGroup.GET("/", s.handler.User.GetUsers)
	userGroup.GET("/{user_id}", s.handler.User.GetByUserID)
	userGroup.PUT("/{user_id}", s.handler.User.PutByUserID)
	userGroup.GET("/{user_id}/posts", s.handler.User.GetUserPosts)
	userGroup.GET("/{user_id}/reactions", s.handler.User.GetUserReactions)
	userGroup.GET("/{user_id}/comments", s.handler.User.GetUserComments)

	postGroup := authGroup.Group("/posts")
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
