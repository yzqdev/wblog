package models

import "time"

func CreateAdmin() {
	var defaultAdmin = &User{
		BaseModel:     BaseModel{},
		Uid:           "thisisuid",
		Username:      "admin",
		Email:         "1@qq.com",
		Telephone:     "",
		Password:      "a66abb5684c45962d887564f08346e8d",
		VerifyState:   "0",
		SecretKey:     "",
		OutTime:       time.Now(),
		GithubLoginId: "",
		GithubUrl:     "",
		IsAdmin:       true,
		AvatarUrl:     "",
		Nickname:      "admin",
		LockState:     false,
	}
	defaultAdmin.Insert()

}
