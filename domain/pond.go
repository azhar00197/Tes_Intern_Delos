package domain

import "gorm.io/gorm"

type PondModel struct {
	gorm.Model
	FarmEntityId uint
	Commodity    string
	Capacity     int
	Price        int64
}

type PondInput struct {
	Commodity string `json:"commodity" binding:"required"`
	Capacity  int    `json:"capacity" binding:"required"`
	Price     int64  `json:"price" binding:"required"`
}

type PondUseCase interface {
	Create(*PondModel) error
	Read(*PondModel) (*[]PondModel, error)
	Update(*PondModel) error
	Delete(*PondModel) error
}
