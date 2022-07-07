package models

import (
	_ "github.com/jinzhu/gorm"
)

type UserSms struct {
	Id        int
	Ip        string
	Phone     string
	SendCount int
	AddDay    string
	AddTime   int
	Sign      string
}

func (UserSms) TableName() string {
	return "user_sms"
}
