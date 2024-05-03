package main

// I think CreatedAt and UpdatedAt fields is default for production ready applications
// But in this test task I don't need them
type Citizen struct {
	Name  string `json:"name" gorm:"not null"`
	Iin   string `json:"iin" gorm:"unique;not null"`
	Phone string `json:"phone" gorm:"not null"`
	// CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	// UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
