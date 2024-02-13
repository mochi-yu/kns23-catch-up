package usecase

import (
	"github.com/mochi-yu/kns23-catch-up/app/model"
	"github.com/mochi-yu/kns23-catch-up/app/repository"
)

type UserUseCase interface {
	RegisterPost(model.RegisterPostParam) error
}

type userUseCase struct {
	r *repository.Repository
}

func NewUserUseCase(r *repository.Repository) UserUseCase {
	return &userUseCase{r: r}
}

// ユーザ登録のエンドポイント
func (uu *userUseCase) RegisterPost(param model.RegisterPostParam) error {
	// リクエストパラメータからモデルを生成
	u := model.RegisterParam2UserModel(param)

	// 新しいユーザをDBに保存
	if err := uu.r.User.InsertNewUser(*u); err != nil {
		return err
	}

	return nil
}
