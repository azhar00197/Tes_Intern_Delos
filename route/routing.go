package route

import (
	_farmhandler "github.com/adityarizkyramadhan/tes_intern_delos/implementation/farm/handler"
	_farmservice "github.com/adityarizkyramadhan/tes_intern_delos/implementation/farm/service"
	"github.com/adityarizkyramadhan/tes_intern_delos/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Route(e *gin.Engine, db *gorm.DB) {
	farmUseCase := _farmservice.NewServiceFarm(db)
	farmHandler := _farmhandler.NewHandlerFarm(farmUseCase)
	farm := e.Group("/farm")
	farm.POST("/", farmHandler.Register)
	farm.PUT("/", middleware.ValidateJWToken(), farmHandler.Update)
	farm.DELETE("/", middleware.ValidateJWToken(), farmHandler.Delete)
	farm.GET("/", middleware.ValidateJWToken(), farmHandler.Read)

}
