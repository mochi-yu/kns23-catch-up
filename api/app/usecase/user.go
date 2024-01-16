package usecase

import "github.com/mochi-yu/kns23-catch-up/app/repository"

type UserUseCase interface {
}

type userUseCase struct {
	r *repository.Repository
}

func NewUserUseCase(r *repository.Repository) UserUseCase {
	return &userUseCase{r: r}
}
