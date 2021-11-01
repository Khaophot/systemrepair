package main

import (
	"github.com/b6226442/systemrepair/controller"

	"github.com/b6226442/systemrepair/entity"

	"github.com/b6226442/systemrepair/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()
	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{
			// Customer Routes
			protected.GET("/customer/:id", controller.GetCustomer)
			// Check In Routes
			protected.GET("/check_in/reserved/:id", controller.ListCheckInsReservedByCustomer)
			// Equipment Routes
			protected.GET("/equipments", controller.ListEquipments)
			// Problem Routes
			protected.GET("/problems", controller.ListProblems)
			// Repairinformation Routes
			protected.GET("/repairinformations", controller.ListRepairInformations)
			protected.POST("/repairinformations", controller.CreateRepairInformation)
			// Urgency Routes
			protected.GET("/urgencies", controller.ListUrgenies)

		}
	}

	// Authentication Routes
	r.POST("/login", controller.Login)

	// Run the server
	r.Run()

}

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
