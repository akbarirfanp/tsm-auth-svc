package models

import (
	"github.com/goravel/framework/database/orm"
)

type User struct {
	orm.Model
	Username     string
	Email    string
	Password string
}
