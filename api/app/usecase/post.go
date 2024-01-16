package usecase

import "github.com/mochi-yu/kns23-catch-up/app/repository"

type PostUseCase interface {
}

type postUseCase struct {
	r *repository.Repository
}

func NewPostUseCase(r *repository.Repository) PostUseCase {
	return &postUseCase{r: r}
}
