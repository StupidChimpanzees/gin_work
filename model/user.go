package model

import (
	"gin_work/wrap/database"
	"time"
)

type User struct {
	Id            int ``
	Uuid          string
	Username      string
	Password      string
	Salt          string
	Nickname      string
	Realname      string
	Email         string
	SignUpIp      string
	CreateTime    time.Duration
	LastLoginIp   string
	LastLoginTime time.Duration
	Avatar        string
	SignUpType    int
	State         int
	DeleteTime    time.Duration
}

func (*User) TableName() string {
	return "y_user"
}

func (u *User) FindById(id int) *User {
	database.DB.Find(&u, id)
	return u
}

func (u *User) FindByUsername(username string) *User {
	database.DB.Where("username = ?", username).First(&u)
	return u
}

func (u *User) FindByEmail(email string) *User {
	database.DB.Where("email = ?", email).First(&u)
	return u
}
