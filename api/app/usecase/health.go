package usecase

import "github.com/mochi-yu/kns23-catch-up/app/repository"

type HealthUseCase interface {
}

type healthUseCase struct {
	r *repository.Repository
}

func NewHealthUseCase(r *repository.Repository) HealthUseCase {
	return &healthUseCase{r: r}
}
