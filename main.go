package main

import (
	"hamedanIND/handler"
	"hamedanIND/util"

	"github.com/gin-gonic/gin"
)

func main() {

	util.Connect()
	db := util.GetDB()

	r := gin.Default()

	apiRoutes := r.Group("/api")
	{
		apiRoutes.POST("/login", handler.Login(&db))
		apiRoutes.POST("/logout", handler.Logout())
		apiRoutes.POST("/capacities", handler.FindCapacities(&db))
		apiRoutes.POST("/industries", handler.FindIndustries(&db))

	}

	r.Run("127.0.0.1:8080")
	defer db.Close()
}
