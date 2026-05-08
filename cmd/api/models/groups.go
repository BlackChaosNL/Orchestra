package models

import (
	"github.com/google/uuid"
)

type Setting struct {
	Base
	Admin  bool `gorm:"type:bool;default:false;"`
	Manage bool `gorm:"type:bool;default:false;"`
	Access bool `gorm:"type:bool;default:true;"`
}

type Membership struct {
	Base
	User    uuid.UUID `gorm:"primaryKey"`
	Group   uuid.UUID `gorm:"primaryKey"`
	Setting Setting
}

type Group struct {
	Base
	Name  string
	Users []User `gorm:"many2many:user_groups"`
}
