package controller

import (
	"gin_work/model"
	"time"
)

func UserLoginUpdate(userInfo *model.User, ip string) {
	user := model.User{
		Id:            userInfo.Id,
		LastLoginIp:   ip,
		LastLoginTime: time.Now().Unix(),
	}
	user.UpdateByLogin()
}
