package models

import (
	"time"
)

type Autoreply struct {
	Id_autoreply int64     `gorm:"primaryKey" json:"id_autoreply"`
	Keyword      string    `gorm:"type:text" json:"id_chat"`
	Answer       string    `gorm:"type:text" json:"name"`
	Created_at   time.Time `gorm:"type:datetime" json:"created"`
	Update_at    time.Time `gorm:"type:datetime" json:"update"`
}
