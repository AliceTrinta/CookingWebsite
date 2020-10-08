package handler

import (
	"fmt"
	"github.com/AliceTrinta/CookingWebsite/lib/auth"
	"github.com/AliceTrinta/CookingWebsite/lib/contx"
	"github.com/AliceTrinta/CookingWebsite/model"
	"github.com/AliceTrinta/CookingWebsite/repo"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strconv"
)

func Signin(ctx *contx.Context)  {
	ctx.NativeHTML(http.StatusOK, "signin")
}

func InsertUser(ctx *contx.Context, form model.User) {
	user, err := repo.GetUserByEmail(form.Email)
	if user == nil {
		if err != nil {
			log.Println("GetUser - repo.GetUser - Error: ", err.Error())
			panic(err)
			return
		}
		err = nil
		_, err := repo.InsertUser(form)
		if err != nil {
			log.Println("InsertUser - repo.InsertUser - Error: ", err.Error())
			panic(err)
			return
		}
		ctx.Redirect("/login")
		return
	}
	ctx.Redirect("/signin")
}

func Login(ctx *contx.Context, user model.User) {
	ans, _ := repo.GetUserByEmailAndPassword(user.Email, user.Password)
	if ans.ID == 0 {
		ctx.Redirect("/login")
		return
	}
	log.Printf("Login - user found!")
	err := auth.CreateJWTCookie(strconv.Itoa(ans.ID), "RecipeSite", 360, ctx)
	if err != nil {
		log.Println("CheckFormUserCredentials - error generating JWT Cookie - Error: ", err.Error())
		ctx.NativeHTML(http.StatusInternalServerError, "erro")
		return
	}
	ctx.Redirect("/")
}

func ShowProfile(ctx *contx.Context)  {
	cookie, _ := ctx.Req.Cookie("RecipeSiteCookie")
	c := cookie.Value
	log.Println(c)
	tokenString := c
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("k3hz0e8ze2pbjn0lg4nlx22at00iwoum"), nil
	})
	if err != nil {
		log.Println("Jwt parse error: ", err.Error())
		panic(err)
		return
	}
	log.Println(token)
	intUserID := 0
	for key, val := range claims {
		fmt.Printf("Key: %v, value: %v\n", key, val)
		if key == "jti"{
			intUserID, _ = strconv.Atoi(fmt.Sprint(val))
		}
	}
	log.Println(intUserID)
	user, err := repo.GetUserByID(intUserID)
	if err != nil {
		log.Println("GetUser - repo.GetUserByID - Error: ", err.Error())
		panic(err)
		return
	}
	ctx.Data["user"] = user
	recipes, err := repo.GetRecipesByUserID(intUserID)
	if err != nil {
		log.Println("GetRecipeByUserID - repo.GetRecipeByUserID - Error: ", err.Error())
		panic(err)
		return
	}
	ctx.Data["recipes"] = recipes
	ctx.NativeHTML(http.StatusOK, "profile")
}