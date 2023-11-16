package handler

import (
	"hamedanIND/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type IndustryNameInfo struct {
	Name       string
}
type IndustryRequest struct {
	Name         string `json:"name"`
}

func FindIndustries(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !util.IsLogedIn(c) {
			c.Redirect(http.StatusFound, "/view/login")
		}
		_, err := c.Cookie("token")
		if err == nil {
			util.BackToLogin(c)
		}
		capacity := IndustryRequest{}
		c.Bind(&capacity)

		industries := []IndustryNameInfo{}
		db.Table("industries as ind").
			Select("name").
			Joins("INNER join capacities as ca on ind.capacity_refer=ca.id").
			Where("ca.name = ?",capacity).
			Scan(&industries)
		data := gin.H{
			"industries":         industries,
		}
		c.JSON(http.StatusOK, data)
	}
}