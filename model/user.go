package model

import "github.com/dgrijalva/jwt-go"

type User struct {
	ID 	 	 	int    `json:"id" db:"id"`
	Name 	 	string `json:"name" db:"name" form:"name"`
	Email 	 	string `json:"email" db:"email" form:"email"`
	Password    string `json:"password" db:"password" form:"password"`
}

type Claims struct {
	IP string `json:"ip"`
	jwt.StandardClaims
}