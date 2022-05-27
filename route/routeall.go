package route

import (
	_farmroute "github.com/adityarizkyramadhan/tes_intern_delos/implementation/farm/delivery_route"
	_farmhandler "github.com/adityarizkyramadhan/tes_intern_delos/implementation/farm/handler"
	_farmservice "github.com/adityarizkyramadhan/tes_intern_delos/implementation/farm/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Route(e *gin.Engine, db *gorm.DB) {
	farmUseCase := _farmservice.NewServiceFarm(db)
	farmHandler := _farmhandler.NewHandlerFarm(farmUseCase)
	farm := e.Group("/farm")
	_farmroute.RouteFarm(farm, farmHandler)
}
