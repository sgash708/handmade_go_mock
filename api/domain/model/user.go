package model

// User 個人情報
type User struct {
	ID      int
	Name    string
	Address string
}

// UserGroup ユーザグループ設定
type UserGroup struct {
	ID      int
	Name    string
	Private bool
}
