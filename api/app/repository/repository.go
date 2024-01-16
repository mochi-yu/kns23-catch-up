package repository

import "github.com/mochi-yu/kns23-catch-up/app/infrastructure"

type Repository struct {
	S3Client infrastructure.S3Client
	User     UserRepository
	Post     PostRepository
}
