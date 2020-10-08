package handler

import (
	"fmt"
	"github.com/AliceTrinta/CookingWebsite/lib/contx"
	"github.com/AliceTrinta/CookingWebsite/model"
	"github.com/AliceTrinta/CookingWebsite/repo"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strconv"
)

func FindAllRecipes(ctx *contx.Context){
	recipes, err := repo.GetAllRecipes()
	if err != nil {
		log.Println("Find all recipes - Error: ", err.Error())
		ctx.NativeHTML(http.StatusInternalServerError, "erro")
		return
	}
	ctx.Data["recipes"] = recipes
	ctx.NativeHTML(http.StatusOK, "allRecipes")
}

func AddRecipe(ctx *contx.Context)  {
	ctx.NativeHTML(http.StatusOK, "addRecipe")
}

func InsertRecipe(ctx *contx.Context, form model.Recipe){
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
	_, err = repo.InsertRecipe(form, intUserID)
	if err != nil {
		log.Println("InsertRecipe - repo.InsertRecipe - Error: ", err.Error())
		panic(err)
		return
	}
	ctx.Redirect("/addRecipe")
}