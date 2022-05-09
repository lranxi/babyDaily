package model

import (
	"encoding/json"
	"time"
)

type Feedback struct {
	ID        int       `json:"id" gorm:"autoIncrement;primaryKey;type:bigint(20)"`
	UserId    int       `json:"userId" gorm:"type:bigint(20);not null"`
	Content   string    `json:"content" gorm:"type:varchar(200);not null"`
	CreatedAt time.Time `json:"createTime" gorm:"type:datetime(0)"`
	UpdatedAt time.Time `json:"updateTime" gorm:"type:datetime(0)"`
}

func (*Feedback) TableName() string {
	return "t_feedback"
}

func (u *Feedback) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func (u *Feedback) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}
