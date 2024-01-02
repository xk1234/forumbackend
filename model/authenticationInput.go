package model

type AuthenticationInput struct {
    Username string `json:"username" binding:"required"`
}