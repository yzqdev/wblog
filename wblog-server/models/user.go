package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

// table users
type User struct {
	gorm.Model
	Email         string    `gorm:"unique_index;default:null"` //邮箱
	Telephone     string    `gorm:"unique_index;default:null"` //手机号码
	Password      string    `gorm:"default:null"`              //密码
	VerifyState   string    `gorm:"default:'0'"`               //邮箱验证状态
	SecretKey     string    `gorm:"default:null"`              //密钥
	OutTime       time.Time //过期时间
	GithubLoginId string    `gorm:"unique_index;default:null"` // github唯一标识
	GithubUrl     string    //github地址
	IsAdmin       bool      //是否是管理员
	AvatarUrl     string    // 头像链接
	NickName      string    // 昵称
	LockState     bool      `gorm:"default:true"` //锁定状态
}

// Insert user User

func (user *User) Insert() error {
	fmt.Println(user)
	return DB.Create(&user).Error
}

// Update update user
func (user *User) Update() error {
	return DB.Save(user).Error
}

//
func GetUserByUsername(username string) (*User, error) {
	var user User
	err := DB.First(&user, "email = ?", username).Error
	return &user, err
}

//
func (user *User) FirstOrCreate() (*User, error) {
	err := DB.FirstOrCreate(user, "github_login_id = ?", user.GithubLoginId).Error
	return user, err
}

func IsGithubIdExists(githubId string, id uint) (*User, error) {
	var user User
	err := DB.First(&user, "github_login_id = ? and id != ?", githubId, id).Error
	return &user, err
}

func GetUser(id interface{}) (*User, error) {
	var user User
	err := DB.First(&user, id).Error
	return &user, err
}

func (user *User) UpdateProfile(avatarUrl, nickName string) error {
	return DB.Model(user).Updates(User{AvatarUrl: avatarUrl, NickName: nickName}).Error
}

func (user *User) UpdateEmail(email string) error {
	if len(email) > 0 {
		return DB.Model(user).Update("email", email).Error
	} else {
		return DB.Model(user).Update("email", gorm.Expr("NULL")).Error
	}
}

func (user *User) UpdateGithubUserInfo() error {
	var githubLoginId interface{}
	if len(user.GithubLoginId) == 0 {
		githubLoginId = gorm.Expr("NULL")
	} else {
		githubLoginId = user.GithubLoginId
	}
	return DB.Model(user).Updates(map[string]interface{}{
		"github_login_id": githubLoginId,
		"avatar_url":      user.AvatarUrl,
		"github_url":      user.GithubUrl,
	}).Error
}

func (user *User) Lock() error {
	return DB.Model(user).Updates(map[string]interface{}{
		"lock_state": user.LockState,
	}).Error
}

func ListUsers() ([]*User, error) {
	var users []*User
	err := DB.Find(&users, "is_admin = ?", false).Error
	return users, err
}
