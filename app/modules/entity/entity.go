package entity

import "time"

type User struct {
	Id         int       `json:"id"`
	Nickname   string    `json:"nickname"`
	DateCreate time.Time `json:"date_create"`
}

type Message struct {
	SendlerId int       `json:"sendler_id"`
	Sendler   User      `json:"sendler" gorm:"foreignKey:Id;references:SendlerId"`
	Text      string    `json:"text"`
	Time      time.Time `json:"time"`
}

type Response struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

const (
	ResponseTypeError = "error"
	ResponseTypeData  = "data"
)
