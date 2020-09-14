package handler

import (
	"github.com/AliceTrinta/cooking-website/conf"
	"github.com/AliceTrinta/cooking-website/lib/contx"
	"net/http"
)

//Index open the initial screen of the webapp
func Index(ctx *contx.Context) {
	db, err := conf.GetDB()
	if err != nil {
		panic(err)
	}
	db.Exec("create database if not exists web-example")
	db.Exec("use web-example")
	db.Exec("create table if not exists user(id integer auto_increment, name varchar(80), email varchar(80), password varchar(80), PRIMARY KEY (id))")
	db.Exec("use web-example")
	db.Exec("create table if not exists recipe(id integer auto_increment, name varchar(80), type varchar(80), description varchar(10000), userId integer, PRIMARY KEY (id), FOREIGN KEY (userId) REFERENCES user(id))")
	ctx.NativeHTML(http.StatusOK, "index")
}
