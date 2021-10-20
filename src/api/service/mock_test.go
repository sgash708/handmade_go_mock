package service_test

import (
	"github.com/sgash708/handmade_go_mock/domain"
	"github.com/sgash708/handmade_go_mock/domain/model"
)

// REF: https://moneyforward.com/engineers_blog/2021/03/08/go-test-mock/
// uRepoMock ユーザリポジトリをモックする構造体
type uRepoMock struct {
	domain.User
	FakeCreate func(user *model.User) (*model.User, error)
}

// Create ユーザ作成モック
func (m *uRepoMock) Create(user *model.User) (*model.User, error) {
	return m.FakeCreate(user)
}

/*
	モックの肝：
	Fake*関数のフィールドをstructに追加する
	→ structのメソッドに求められるメソッドを実装

	※メソッドの定義がインターフェースと異なるとエラーを吐く
*/
// ugRepoMock ユーザグループリポジトリをモックする構造体
type ugRepoMock struct {
	domain.UserGroup
	FakeCreate  func(ug *model.UserGroup) (*model.UserGroup, error)
	FakeAddUser func(ugID int, uID int) error
}

// Create ユーザグループ作成モック
func (m *ugRepoMock) Create(ug *model.UserGroup) (*model.UserGroup, error) {
	return m.FakeCreate(ug)
}

// AddUser ユーザグループ追加モック
func (m *ugRepoMock) AddUser(ugID int, uID int) error {
	return m.FakeAddUser(ugID, uID)
}

/*
	この状態でテストできる理由：
	struct(uRepoMock)にinterface(domain.User)を埋め込む
	→ interfaceは匿名フィールドになる
	→ interfaceは必要なメソッド定義を全て持っている
		→ structも必要なメソッド定義を持っている
		→ interfaceを満たしている
*/
