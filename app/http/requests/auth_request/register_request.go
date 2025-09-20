package auth_request

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type RegisterRequest struct {
	Username string `form:"username" json:"username"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

func (r *RegisterRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *RegisterRequest) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *RegisterRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"username": "required|max_len:50",
		"email":    "required|max_len:70",
		"password": "required|max_len:255",
	}
}

func (r *RegisterRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"username.required": "Username is required",
		"email.required":    "Email address is required",
		"password.required": "Password is required",
	}
}

func (r *RegisterRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *RegisterRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
