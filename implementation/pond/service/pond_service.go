package service

import (
	"github.com/adityarizkyramadhan/tes_intern_delos/domain"
	"gorm.io/gorm"
)

type servicePond struct {
	db *gorm.DB
}

func NewServicePond(db *gorm.DB) domain.PondUseCase {
	return &servicePond{
		db: db,
	}
}

func (s servicePond) Create(model *domain.PondModel) error {
	//TODO implement me
	panic("implement me")
}

func (s servicePond) Read(u uint) (*[]domain.PondModel, error) {
	//TODO implement me
	panic("implement me")
}

func (s servicePond) Update(u uint) error {
	//TODO implement me
	panic("implement me")
}

func (s servicePond) Delete(u uint) error {
	//TODO implement me
	panic("implement me")
}
