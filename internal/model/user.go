package model

import "time"

type User struct {
	Id 			int64		`json:"id" gorm:"id"`
	Name 		string		`json:"name" gorm:"name"`
	Age 		int32		`json:"age" gorm:"age"`
	Gender 		int32		`json:"gender" gorm:"gender"`
	Email 		string		`json:"email" gorm:"email"`
	CreateAt 	time.Time	`json:"createAt" gorm:"create_at"`
	UpdateAt 	time.Time	`json:"updateAt" gorm:"update_at"`
}

func (*User) TableName() string {
	return "user"
}
