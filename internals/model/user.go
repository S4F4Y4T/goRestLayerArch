package model

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"uniqueIndex"`
}

type UserRepository interface {
	Create(user *User) error
	FindAll() ([]User, error)
	FindByID(id uint) (*User, error)
}
