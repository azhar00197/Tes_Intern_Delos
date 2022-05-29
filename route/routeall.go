package route

import (
	_farmRoute "github.com/adityarizkyramadhan/tes_intern_delos/implementation/farm/delivery_route"
	_farmHandler "github.com/adityarizkyramadhan/tes_intern_delos/implementation/farm/handler"
	_farmService "github.com/adityarizkyramadhan/tes_intern_delos/implementation/farm/service"
	_pondRoute "github.com/adityarizkyramadhan/tes_intern_delos/implementation/pond/delivery_route"
	_pondHandler "github.com/adityarizkyramadhan/tes_intern_delos/implementation/pond/handler"
	_pondService "github.com/adityarizkyramadhan/tes_intern_delos/implementation/pond/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Route(e *gin.Engine, db *gorm.DB) {
	farmUseCase := _farmService.NewServiceFarm(db)
	farmHandler := _farmHandler.NewHandlerFarm(farmUseCase)
	farm := e.Group("/farm")
	_farmRoute.RouteFarm(farm, farmHandler)
	pondUseCase := _pondService.NewServicePond(db)
	pondHandler := _pondHandler.NewHandlerPond(pondUseCase)
	pond := e.Group("/pond")
	_pondRoute.RoutePonds(pond, pondHandler)
}
