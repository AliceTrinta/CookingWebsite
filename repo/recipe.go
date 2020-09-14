package repo

import (
	"github.com/AliceTrinta/cooking-website/conf"
	"github.com/AliceTrinta/cooking-website/model"
)

//Getting all recipes stored on database
func GetAllRecipes() (recipes []model.Recipe, err error){
	db, err := conf.GetDB()
	if err != nil {
		return
	}
	err = db.Select(&recipes, "SELECT * FROM recipe")
	if err != nil {
		return
	}
	return
}
