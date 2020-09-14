package handler

import (
	"github.com/AliceTrinta/cooking-website/lib/contx"
	"github.com/AliceTrinta/cooking-website/repo"
	"log"
	"net/http"
)

//Func to find all recipes stored on database
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
