package models

type User struct {
	Base
	Username string  `gorm:"uniqueIndex;not null" json:"username"`
	Email    string  `gorm:"uniqueIndex;not null" json:"email"`
	Password string  `gorm:"not null" json:"-"`
	Groups   []Group `gorm:"many2many:user_groups"`
}
