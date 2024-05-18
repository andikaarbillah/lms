package model

import "time"

type Users struct {
	Id         string    `json:"id" gorm:"primaryKey;type:varchar(55);index"`
	Username   string    `json:"username" gorm:"type:varchar(55);not null"`
	Email      string    `json:"email" validate:"required,email" gorm:"type:varchar(55);not null;unique"`
	Password   string    `json:"password" validate:"required" gorm:"type:varchar(255)"`
	Created_at time.Time `json:"created_at" gorm:"type:timestamp default current_timestamp"`
	Updated_at time.Time `json:"updated_at" gorm:"type:timestamp;not null"`
	Is_deleted bool      `json:"is_deleted" gorm:"type:boolean default false"`
}
