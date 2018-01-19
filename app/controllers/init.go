package controllers

import (
	"github.com/revel/revel"
)

func InitDB() {
	revel.INFO.Println("DB Connected")
}

func init() {
	revel.OnAppStart(InitDB)
}
