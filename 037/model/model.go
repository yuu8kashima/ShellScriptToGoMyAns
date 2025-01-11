package model

//user.go
type User struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	PasswordHash string `gorm:"not null" json:"password_hash,omitempty"`
}
