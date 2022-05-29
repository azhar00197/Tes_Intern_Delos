package delivery_route

import (
	"github.com/adityarizkyramadhan/tes_intern_delos/implementation/farm/handler"
	"github.com/adityarizkyramadhan/tes_intern_delos/middleware"
	"github.com/gin-gonic/gin"
)

func RouteFarm(farm *gin.RouterGroup, farmHandler *handler.FarmHandler) {
	farm.POST("/", farmHandler.FarmHandler.Create)
	farm.PUT("/", middleware.ValidateJWToken(), farmHandler.FarmHandler.Update)
	farm.DELETE("/", middleware.ValidateJWToken(), farmHandler.FarmHandler.Delete)
	farm.GET("/", middleware.ValidateJWToken(), farmHandler.FarmHandler.Read)
}
