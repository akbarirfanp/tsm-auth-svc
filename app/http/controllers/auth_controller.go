package controllers

import (
	"time"

	"goravel/app/http/requests/auth_request"
	"goravel/app/models"
	"goravel/app/traits"

	"github.com/golang-jwt/jwt/v5"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/validation"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	response traits.ResponseAPI
}

func NewAuthController() *AuthController {
	return &AuthController{
		response: traits.ResponseAPI{},
	}
}

// Register new user
func (r *AuthController) Register(ctx http.Context) http.Response {
	request := &auth_request.RegisterRequest{}

	// Bind request
	if err := ctx.Request().Bind(request); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"error": "Invalid request body",
		})
	}

	// Validate
	validator, err := facades.Validation().Make(
		ctx.Request().All(),
		request.Rules(ctx),
		validation.Messages(request.Messages(ctx)),
	)
	if err != nil {
		facades.Log().Errorf("Validation make error: %v", err)
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": "Validation service error",
		})
	}
	if validator.Fails() {
		return ctx.Response().Json(http.StatusUnprocessableEntity, http.Json{
			"errors": validator.Errors().All(),
		})
	}

	// hash password
	hashed, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	user := models.User{
		Username: request.Username,
		Email:    request.Email,
		Password: string(hashed),
	}

	if err := facades.Orm().Query().Create(&user); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"error": err.Error(),
		})
	}

	return r.response.Success(ctx, nil, "User registered successfully")
}

// Login user
func (r *AuthController) Login(ctx http.Context) http.Response {
	request := &auth_request.LoginRequest{}

	if err := ctx.Request().Bind(request); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"error": "Invalid request body",
		})
	}

	validator, err := facades.Validation().Make(
		ctx.Request().All(),
		request.Rules(ctx),
		validation.Messages(request.Messages(ctx)),
	)
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
	}
	if validator.Fails() {
		return ctx.Response().Json(http.StatusUnprocessableEntity, http.Json{
			"errors": validator.Errors().All(),
		})
	}

	var user models.User
	if err := facades.Orm().Query().Where("email", request.Email).First(&user); err != nil {
		return ctx.Response().Json(http.StatusUnauthorized, http.Json{
			"error": "Invalid credentials",
		})
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		return ctx.Response().Json(http.StatusUnauthorized, http.Json{
			"error": "Invalid credentials",
		})
	}

	secret := []byte(facades.Config().GetString("jwt.secret"))
	ttl := facades.Config().GetInt("jwt.ttl")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Minute * time.Duration(ttl)).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": "Failed to generate token",
		})
	}

	// hanya email + token
	data := map[string]any{
		"email": user.Email,
		"token": tokenString,
	}

	return r.response.Success(ctx, data, "Login successful")
}
