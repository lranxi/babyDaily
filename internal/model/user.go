package model

import (
	"encoding/json"
	"time"
)

type User struct {
	ID        int       `json:"id" gorm:"autoIncrement;primaryKey;type:bigint(20)"`
	Nickname  string    `json:"name" gorm:"type:varchar(50);not null"`
	OpenId    string    `json:"openId" gorm:"type:varchar(50);unique;not null"`
	UnionId   string    `json:"email" gorm:"type:varchar(50);not null"`
	Avatar    string    `json:"avatar" gorm:"type:varchar(200)"`
	CreatedAt time.Time `json:"createTime" gorm:"type:datetime(0)"`
	UpdatedAt time.Time `json:"updateTime" gorm:"type:datetime(0)"`
}

func (*User) TableName() string {
	return "t_user"
}

func (u *User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func (u *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}

func (u *User) CacheKey() string {
	return u.TableName() + ":openId"
}

type CreatedUser struct {
	Code     string `json:"code"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type UpdatedUser struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

func (u *UpdatedUser) GetUser() *User {
	return &User{
		Nickname:  u.Nickname,
		Avatar:    u.Avatar,
		UpdatedAt: time.Now(),
	}
}

// AuthUser 授权用户信息
type AuthUser struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
	OpenId   string `json:"openId"`
	Avatar   string `json:"avatar"`
	Token    string `json:"token"`
	Expires  uint64 `json:"expires"`
}
