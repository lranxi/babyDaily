package model

import (
	"encoding/json"
	"time"
)

type Babies []Baby

type Baby struct {
	ID        int       `json:"id" gorm:"autoIncrement;primaryKey;type:bigint(20)"`
	UserId    int       `json:"userId" gorm:"type:bigint(20);not null;index"`
	Name      string    `json:"name" gorm:"type:varchar(6);default: ''"`
	Height    int       `json:"height" gorm:"type:int(4)"`
	Weight    int       `json:"weight" gorm:"type:int(4)"`
	Brith     string    `json:"brith" gorm:"type:varchar(20);"`
	Gender    uint8     `json:"gender" gorm:"type:tinyint(1)"`
	Photo     string    `json:"photo" gorm:"type:varchar(200)"`
	CreatedAt time.Time `json:"createTime" gorm:"type:datetime(0)"`
	UpdatedAt time.Time `json:"updateTime" gorm:"type:datetime(0)"`
}

func (b *Baby) TableName() string {
	return "t_baby"
}

func (b *Baby) MarshalBinary() ([]byte, error) {
	return json.Marshal(b)
}

func (b *Baby) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, b)
}

func (b *Baby) CacheKey() string {
	return b.TableName() + ":id"
}

// CreateBaby 新增baby结构体
type CreateBaby struct {
	UserId int    `json:"userId"`
	Name   string `json:"name"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Brith  string `json:"brith"`
	Gender uint8  `json:"gender"`
	Photo  string `json:"photo"`
}

// GetBaby 新增baby转换为baby结构体
func (b *CreateBaby) GetBaby() *Baby {
	return &Baby{
		UserId:    b.UserId,
		Name:      b.Name,
		Height:    b.Height,
		Weight:    b.Weight,
		Brith:     b.Brith,
		Gender:    b.Gender,
		Photo:     b.Photo,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
