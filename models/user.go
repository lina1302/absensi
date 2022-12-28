package models

import (
	"time"
)

type User struct {
	Id_user    int64     `gorm:"primaryKey" json:"id_user"`
	Id_chat    int64     `gorm:"type:int" json:"id_chat"`
	Type       string    `gorm:"enum" json:"type"`
	Name       string    `gorm:"type:varchar(100)" json:"name"`
	Username   string    `gorm:"type:varchar(50)" json:"username"`
	Created_at time.Time `gorm:"type:datetime" json:"created"`
	Update_at  time.Time `gorm:"type:datetime" json:"update"`
}
