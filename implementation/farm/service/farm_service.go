package service

import (
	"github.com/adityarizkyramadhan/tes_intern_delos/domain"
	"gorm.io/gorm"
)

type serviceFarm struct {
	db *gorm.DB
}

func NewServiceFarm(db *gorm.DB) domain.FarmUseCase {
	return &serviceFarm{
		db: db,
	}
}

func (s serviceFarm) Create(model *domain.FarmModel) (uint, error) {
	if err := s.db.Model(domain.FarmModel{}).Create(&model).Error; err != nil {
		return 0, err
	}
	return model.ID, nil
}

func (s serviceFarm) Read(id uint) (domain.FarmModel, error) {
	var readFarm domain.FarmModel
	if err := s.db.Model(&domain.FarmModel{}).Where("id = ?", id).Preload("PondModel").Take(&readFarm).Error; err != nil {
		return domain.FarmModel{}, err
	}
	return readFarm, nil
}

func (s serviceFarm) Delete(id uint) error {
	if err := s.db.Model(&domain.FarmModel{}).Delete("id = ?", id).Error; err != nil {
		return err
	}
	if err := s.db.Model(&domain.PondModel{}).Delete("farm_entity_id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (s serviceFarm) Update(model *domain.FarmModel, id uint) error {
	if err := s.db.Model(&domain.FarmModel{}).Where("id = ?", id).Save(&model).Error; err != nil {
		return err
	}
	return nil
}
