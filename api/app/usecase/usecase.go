package usecase

type UseCase struct {
	Health HealthUseCase
	Post   PostUseCase
	User   UserUseCase
}
