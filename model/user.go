package model

import (
	"gin_work/wrap/database"
)

type User struct {
	Id            uint   `gorm:"primaryKey"`
	Uuid          string `gorm:"uuid,<-:create"`
	Username      string `gorm:"username,<-:create"`
	Password      string `gorm:"password"`
	Salt          string `gorm:"salt"`
	Nickname      string `gorm:"nickname"`
	RealName      string `gorm:"realname"`
	Email         string `gorm:"email"`
	SignUpIp      string `gorm:"sign_up_ip,<-:create"`
	CreateTime    int64  `gorm:"create_time,autoCreateTime"`
	LastLoginIp   string `gorm:"last_login_ip"`
	LastLoginTime int64  `gorm:"last_login_time"`
	Avatar        string `gorm:"avatar"`
	SignUpType    int    `gorm:"sign_up_type,<-:create"`
	State         int    `gorm:"state"`
	DeleteTime    int64  `gorm:"delete_time"`
}

func (*User) TableName() string {
	return "y_user"
}

func (u *User) FindById(id uint) {
	database.DB.Find(&u, id)
}

func (u *User) FindByUuid(uuid string) {
	database.DB.Where("uuid = ?", uuid).First(&u)
}

func (u *User) FindByUsername(username string) {
	database.DB.Where("username = ?", username).First(&u)
}

func (u *User) FindByEmail(email string) *User {
	database.DB.Where("email = ?", email).First(&u)
	return u
}

func (u *User) UpdateByLogin() {
	database.DB.Save(u)
}
