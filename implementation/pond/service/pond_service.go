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
	if err := s.db.Model(&domain.PondModel{}).Create(&model).Error; err != nil {
		return err
	}
	return nil
}

func (s servicePond) Read(id uint) ([]domain.PondModel, error) {
	var data []domain.PondModel
	err := s.db.Model(&domain.PondModel{}).Where("farm_entity_id = ?", id).Find(&data).Error
	if data == nil && err != nil {
		return nil, err
	} else if data == nil && err == nil {
		return nil, nil
	}
	return data, nil
}

func (s servicePond) Update(id uint, model *domain.PondModel) error {
	if err := s.db.Model(&domain.PondModel{}).Where("id = ?", id).Save(&model).Error; err != nil {
		return err
	}
	return nil
}

func (s servicePond) Delete(id uint) error {
	if err := s.db.Model(&domain.PondModel{}).Delete("id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (s servicePond) ReadById(id uint) (domain.PondModel, error) {
	var data domain.PondModel
	err := s.db.Model(&domain.PondModel{}).Where("id = ?", id).Take(&data).Error
	if data == (domain.PondModel{}) && err != nil {
		return data, err
	}
	return data, nil
}
