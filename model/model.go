package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string
	UserName string
	Password string
	Gender   bool
	Birthday time.Time
}

type City struct{
	gorm.Model
	Name string
}

type Capacity struct{
	gorm.Model
	Name 	  string
	City      City `gorm:"foreignkey:CityRefer"`
	CityRefer uint
}

type Industry struct{
	gorm.Model
	Name 	  string
	Capacity      Capacity `gorm:"foreignkey:CapacityRefer"`
	CapacityRefer uint
}
