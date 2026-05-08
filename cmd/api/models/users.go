package models

type User struct {
	Base
	Username string  `gorm:"uniqueIndex;not null" json:"username"`
	Email    string  `gorm:"uniqueIndex;not null" json:"email"`
	Groups   []Group `gorm:"many2many:user_groups"`
}
