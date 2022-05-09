package model

import (
	"encoding/json"
	"time"
)

type Records []Record

// Record 每日记录
type Record struct {
	ID        int       `json:"id" gorm:"autoIncrement;primaryKey;type:bigint(20)"`
	BabyId    int       `json:"babyId" gorm:"type:bigint(20);not null;index"`
	UserId    int       `json:"userId" gorm:"type:bigint(20);not null;index"`
	Type      int       `json:"type" gorm:"type:int(2);not null"`
	SubType   int       `json:"subType" gorm:"type:int(2);not null;default: 0"`
	Quantity  int       `json:"quantity" gorm:"type:int(4);default: 0"`
	Start     string    `json:"start" gorm:"type:varchar(20);not null"`
	End       string    `json:"end" gorm:"type:varchar(20);"`
	Remark    string    `json:"remark" gorm:"type:varchar(200)"`
	CreatedAt time.Time `json:"createTime" gorm:"type:datetime(0)"`
	UpdatedAt time.Time `json:"updateTime" gorm:"type:datetime(0)"`
}

func (*Record) TableName() string {
	return "t_record"
}

func (u *Record) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func (u *Record) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}

func (u *Record) CacheKey() string {
	return u.TableName() + ":openId"
}

type CreateRecord struct {
	BabyId  int    `json:"babyId"`
	UserId  int    `json:"userId"`
	Type    int    `json:"type"`
	SubType int    `json:"subType"`
	Start   string `json:"start"`
	End     string `json:"end"`
	Remark  string `json:"remark"`
}

type Details []RecordDetail

// RecordDetail 记录详情
type RecordDetail struct {
	BabyId    int                `json:"babyId"`
	UserId    int                `json:"userId"`
	Type      int                `json:"type"`
	SubRecord []*SubRecordDetail `json:"subRecord"`
}

type SubRecordDetail struct {
	SubType  int `json:"subType"`
	Quantity int `json:"quantity" gorm:"type:int(4);"`
}

func (c *CreateRecord) GetRecord() *Record {
	return &Record{
		BabyId:    c.BabyId,
		UserId:    c.UserId,
		Type:      c.Type,
		SubType:   c.SubType,
		Start:     c.Start,
		End:       c.End,
		Remark:    c.Remark,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
