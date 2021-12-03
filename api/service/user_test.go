package service_test

import (
	"testing"

	"github.com/sgash708/handmade_go_mock/domain/model"
	"github.com/sgash708/handmade_go_mock/service"
)

func TestUser_Create(t *testing.T) {
	uRepo := &uRepoMock{
		FakeCreate: func(user *model.User) (*model.User, error) {
			created := &model.User{ID: 7, Name: user.Name, Address: user.Address}
			return created, nil
		},
	}
	ugRepo := &ugRepoMock{
		FakeCreate: func(ug *model.UserGroup) (*model.UserGroup, error) {
			created := &model.UserGroup{ID: 9, Name: ug.Name, Private: ug.Private}
			return created, nil
		},
		FakeAddUser: func(ugID int, uID int) error { return nil },
	}

	userService := service.NewUser(uRepo, ugRepo)

	userInput := model.User{Name: "John", Address: "Kyoto"}
	userData, err := userService.Create(&userInput)
	if err != nil {
		t.Fatal(err)
	}

	if userData.ID != 7 {
		t.Errorf("User.Create()は、model.User.ID(7)を返すはずです。\n実際の入力：%d", userData.ID)
	}
	if userData.Name != userInput.Name {
		t.Errorf("User.Create()は、model.User.Name(John)はを返すはずです。\n実際の入力：%s", userData.Name)
	}
	if userData.Address != userInput.Address {
		t.Errorf("User.Create()は、model.User.Address(Kyoto)を返すはずです。\n実際の入力：%s", userData.Address)
	}
}

// func TestUser_Read(t *testing.T) {
// 	uRepo := &uRepoMock{
// 		FakeRead: func(user *model.User) (*model.User, error) {
// 			readed := &model.User{ID: 7, Name: "John", Address: "Kyoto"}
// 			return readed, nil
// 		},
// 	}

// 	userService := service.NewUser(uRepo, nil)

// 	userInput := model.User{Name: "John", Address: "Kyoto"}
// 	userData, err := userService.Read(&userInput)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if userData.ID != 7 {
// 		t.Errorf("User.Create()は、model.User.ID(7)を返すはずです。\n実際の入力：%d", userData.ID)
// 	}
// 	if userData.Name != userInput.Name {
// 		t.Errorf("User.Create()は、model.User.Name(John)はを返すはずです。\n実際の入力：%s", userData.Name)
// 	}
// 	if userData.Address != userInput.Address {
// 		t.Errorf("User.Create()は、model.User.Address(Kyoto)を返すはずです。\n実際の入力：%s", userData.Address)
// 	}
// }
