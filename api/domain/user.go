package domain

import "github.com/sgash708/handmade_go_mock/domain/model"

// User UserモデルへのCRUDを表す
type User interface {
	Read(uID int) (*model.User, error)
	Create(u *model.User) (*model.User, error)
	Update(u *model.User) error
	Delete(uID int) error
}

// UserGroup UserGroupモデルへのCRUDを表す
type UserGroup interface {
	Read(ugID int) (*model.UserGroup, error)
	Create(ug *model.UserGroup) (*model.UserGroup, error)
	Update(ug *model.UserGroup) error
	Delete(ug int) error

	ListUser(ugID int) ([]*model.User, error)
	AddUser(ugID int, uID int) error
	DeleteUser(ugID int, uID int) error
}
