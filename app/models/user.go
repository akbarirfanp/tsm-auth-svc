package models

import (
	"github.com/google/uuid"
	"github.com/goravel/framework/database/orm"
)

type User struct {
	orm.Model
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	Username string
	Email    string
	Password string
}
