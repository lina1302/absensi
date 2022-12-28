package models

import (
	"time"
)

type Chat struct {
	Id_chat      int64     `gorm:"primaryKey" json:"id_chat"`
	Id_user      int64     `gorm:"type:int" json:"id_user"`
	Id_autoreply int64     `gorm:"type:int" json:"id_autoreply"`
	Text         string    `gorm:"type:text" json:"text"`
	Send_at      time.Time `gorm:"type:datetime" json:"send"`
	Created_at   time.Time `gorm:"type:datetime" json:"created"`
	Update_at    time.Time `gorm:"type:datetime" json:"update"`
}
