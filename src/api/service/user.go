package service

import (
	"errors"
	"fmt"

	"github.com/sgash708/handmade_go_mock/domain"
	"github.com/sgash708/handmade_go_mock/domain/model"
)

// User 個人情報の管理
type User struct {
	uRepo  domain.User
	ugRepo domain.UserGroup
}

// NewUser UserServiceの初期化
func NewUser(uRepo domain.User, ugRepo domain.UserGroup) *User {
	return &User{uRepo: uRepo, ugRepo: ugRepo}
}

// Create 登録処理
func (u *User) Create(user *model.User) (*model.User, error) {
	/*
		バリデーション
		TODO: DDDにおけるバリデーションはどの層に記述すべきか考えたい....
	*/
	if user.Name == "" {
		return nil, errors.New("user name required")
	}

	// ユーザ作成
	createdUser, err := u.uRepo.Create(user)
	if err != nil {
		return nil, err
	}

	// グループ作成
	createdUG, err := u.ugRepo.Create(
		&model.UserGroup{
			Name:    fmt.Sprintf("%s's default group", user.Name),
			Private: true,
		},
	)
	if err != nil {
		return nil, err
	}

	// 作成したユーザをグループに追加
	err = u.ugRepo.AddUser(createdUG.ID, createdUser.ID)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
