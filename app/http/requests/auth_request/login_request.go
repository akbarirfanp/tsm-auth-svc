package auth_request

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type LoginRequest struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"passwod" json:"password"`
}

func (r *LoginRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *LoginRequest) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *LoginRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"email":    "required|max_len:70",
		"password": "required|max_len:255",
	}
}

func (r *LoginRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"email.required":    "Email address is required",
		"password.required": "Password is required",
	}
}

func (r *LoginRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *LoginRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
