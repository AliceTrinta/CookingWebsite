package repo

import (
	"github.com/AliceTrinta/CookingWebsite/conf"
	"github.com/AliceTrinta/CookingWebsite/model"
)

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

func GetRecipesByUserID(userId int) (recipes []model.Recipe, err error) {
	db, err := conf.GetDB()
	if err != nil {
		return
	}
	err = db.Select(&recipes, "SELECT * FROM recipe where userId=?", userId)
	if err != nil {
		return
	}
	return
}

func InsertRecipe(form model.Recipe, userId int) (newRecipe model.Recipe, err error)   {
	db, err := conf.GetDB()
	if err != nil {
		return
	}
	println(userId)
	ans, err := db.Exec("INSERT INTO recipe (name, type, description, userId) VALUES (?, ?, ?, ?)", form.Name, form.Type, form.Description, userId)
	if err != nil {
		return
	}
	newID, err := ans.LastInsertId()
	if err != nil {
		return
	}
	newRecipe.ID = int(newID)
	newRecipe.Name  = form.Name
	newRecipe.Type = form.Type
	newRecipe.Description = form.Description
	newRecipe.UserID = userId
	return
}