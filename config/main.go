package main

import (
	"hamedanIND/model"
	"hamedanIND/util"
)

// Database Migration:
func main() {
	util.Connect()
	d := util.GetDB()
	d.AutoMigrate(&model.User{},
		&model.City{},
		&model.Capacity{},
		&model.Industry{})
	defer d.Close()
}
