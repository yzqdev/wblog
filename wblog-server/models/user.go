package models

import (
	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
	"time"
)

// table users
type User struct {
	BaseModel
	Uid           string    `json:"uid" gorm:"type:varchar(64);unique_index;"`
	Username      string    `json:"username" gorm:"type:varchar(64);unique_index;"`
	Email         string    `json:"email" gorm:"type:varchar(64);unique_index;default:null"`         //邮箱
	Telephone     string    `json:"telephone"gorm:"type:varchar(64);unique_index;default:null"`      //手机号码
	Password      string    `json:"-" gorm:"type:varchar(64);default:null"`                          //密码
	VerifyState   string    `json:"verify_state" gorm:"type:varchar(64);default:'0'"`                //邮箱验证状态
	SecretKey     string    `json:"secret_key" gorm:"type:varchar(64);default:null"`                 //密钥
	OutTime       time.Time `json:"out_time"`                                                        //过期时间
	GithubLoginId string    `json:"githubLoginId" gorm:"type:varchar(64);unique_index;default:null"` // github唯一标识
	GithubUrl     string    `json:"github_url" gorm:"type:varchar(64);"`                             //github地址
	AdminRole     bool      `json:"admin_role"`                                                      //是否是管理员
	AvatarUrl     string    `json:"avatar_url" gorm:"type:varchar(64);"`                             // 头像链接
	Nickname      string    `json:"nickname" gorm:"type:varchar(64);"`                               // 昵称
	LockState     bool      `json:"lock_state" gorm:"type:bool;default:false"`                       //锁定状态
}

// user
// insert user
func (user *User) Insert() error {
	node, _ := snowflake.NewNode(1)
	user.Id = node.Generate().String()
	return DB.Create(user).Error
}

// update user
func (user *User) Update() error {
	return DB.Save(user).Error
}

//
func GetUserByUsername(username string) (*User, error) {
	var user User
	err := DB.First(&user, "username = ?", username).Error
	return &user, err
}

//
func (user *User) FirstOrCreate() (*User, error) {
	err := DB.FirstOrCreate(user, "github_login_id = ?", user.GithubLoginId).Error
	return user, err
}

func IsGithubIdExists(githubId string, id string) (*User, error) {
	var user User
	err := DB.First(&user, "github_login_id = ? and id != ?", githubId, id).Error
	return &user, err
}

func GetUser(id interface{}) (*User, error) {
	var user User
	err := DB.First(&user, id).Error
	return &user, err
}
func GetUserByUid(id interface{}) (*User, error) {
	var user User
	err := DB.First(&user, "uid=?", id).Error
	return &user, err
}

func (user *User) UpdateProfile(avatarUrl, nickName string) error {
	return DB.Model(user).Where("uid=?", user.Uid).Updates(User{AvatarUrl: avatarUrl, Nickname: nickName}).Error
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
