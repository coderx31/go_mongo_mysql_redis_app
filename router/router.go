package router

import (
	"mongo/configs"
	"mongo/controller"

	"github.com/gin-gonic/gin"
)

func RouteInitializer() {

	config, _ := configs.Configs()

	r := gin.Default()

	r.GET("/api/people", controller.GetPeople)
	r.GET("/api/people/:firstname", controller.GetPerson)
	r.POST("/api/people", controller.CreatePerson)
	r.PUT("/api/people/:firstname", controller.UpdateCustomerAge)
	r.DELETE("/api/people/:firstname", controller.DeletePerson)

	r.Run(config.App.Port)

}
