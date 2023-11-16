package handler

import (
	"hamedanIND/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CapacityNameInfo struct {
	Name       string
}
type CapacityRequest struct {
	Name         string `json:"name"`
}

func FindCapacities(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !util.IsLogedIn(c) {
			c.Redirect(http.StatusFound, "/view/login")
		}
		_, err := c.Cookie("token")
		if err == nil {
			util.BackToLogin(c)
		}
		
		city := CapacityRequest{}
		c.Bind(&city)

		capacities := []CapacityNameInfo{}
		db.Table("capacities as ca").
			Select("name").
			Joins("INNER join cities as ci on ca.city_refer=ci.id").
			Where("ci.name = ?",city).
			Scan(&capacities)
		data := gin.H{
			"capacities":         capacities,
		}
		c.JSON(http.StatusOK, data)
	}
}