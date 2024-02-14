package usecase

import (
	"firebase.google.com/go/v4/auth"
	"github.com/mochi-yu/kns23-catch-up/app/model"
	"github.com/mochi-yu/kns23-catch-up/app/repository"
)

type UserUseCase interface {
	RegisterPost(model.RegisterPostParam) error
	GetAuth(u *auth.UserInfo) (
		userInfo *model.UserModel,
		is_temp_user bool,
		err error,
	)
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

// ユーザ情報取得のエンドポイント
func (uu *userUseCase) GetAuth(u *auth.UserInfo) (
	userInfo *model.UserModel,
	is_temp_user bool,
	err error,
) {
	// usersテーブルからユーザを取得
	userInfo, err = uu.r.User.GetUserByUid(u.UID)
	if err != nil {
		// ユーザが見つからない場合には"no user"のエラーを返す
		if err.Error() != "no user" {
			return nil, false, err
		}
	} else {
		return userInfo, false, nil
	}

	// 取得に失敗したらtemp_usersテーブルからユーザを取得
	tempUser, err := uu.r.User.GetTempUserByUid(u.UID)
	if err != nil {
		if err.Error() != "no temp_user" {
			return nil, false, err
		}
	} else {
		return model.TempUserModel2UserModel(tempUser), true, nil
	}

	// それもできなければ、新しくtemp_usersテーブルに格納
	tempUser = &model.TempUserModel{
		FirebaseID:  u.UID,
		MailAddress: u.Email,
	}

	err = uu.r.User.InsertTempUser(*tempUser)
	if err != nil {
		return nil, false, err
	}

	return model.TempUserModel2UserModel(tempUser), true, nil
}
