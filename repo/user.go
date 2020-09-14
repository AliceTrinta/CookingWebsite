package repo

import (
	"github.com/AliceTrinta/cooking-website/conf"
	"github.com/AliceTrinta/cooking-website/model"
	"log"
)

func InsertUser(form model.User) (newUser model.User, err error)  {
	db, err := conf.GetDB()
	if err != nil {
		return
	}
	ans, err := db.Exec("INSERT INTO user (name, email, password) VALUES (?, ?, ?)", form.Name, form.Email, form.Password)
	if err != nil {
		return
	}
	newID, err := ans.LastInsertId()
	if err != nil {
		return
	}
	newUser.ID = int(newID)
	newUser.Name  = form.Name
	newUser.Email = form.Email
	newUser.Password = form.Password
	return
}

func GetUserByEmail(email string) (user []model.User, err error){
	db, err := conf.GetDB()
	if err != nil {
		return
	}
	err = db.Select(&user, "SELECT * FROM `web-example`.user WHERE email=?", email)
	if err != nil {
		return
	}
	return
}

func GetUserByID(id int) (user []model.User, err error){
	db, err := conf.GetDB()
	if err != nil {
		return
	}
	err = db.Select(&user, "SELECT * FROM `web-example`.user WHERE id=?", id)
	if err != nil {
		return
	}
	return
}

func GetUserByEmailAndPassword(email string, password string) (user model.User, err error) {
	db, err := conf.GetDB()
	if err != nil {
		log.Panic(err)
		return
	}
	err = db.Get(&user, "SELECT * FROM `web-example`.user WHERE email=? and password=?", email, password)
	if err != nil {
		log.Println("Error db.Select GetUserByCredentials - Erro: ", err.Error())
		return
	}
	return
}