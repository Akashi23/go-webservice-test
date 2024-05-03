package main

import "time"

type Citizen struct {
	Name      string    `json:"name" gorm:"not null"`
	Iin       string    `json:"iin" gorm:"unique;not null"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
