package repository

import (
	"gorm.io/gorm"
)

type BaseRepository struct {
	model *gorm.DB
}

func NewRepository(model *gorm.DB) (Repository, error) {
	return &BaseRepository{
		model: model,
	}, nil
}

func (b *BaseRepository) Find(id int32) (interface{}, error) {
	return nil, nil
}

func (b *BaseRepository) Select(query map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	err := b.model.Where(query).First(&result).Error
	if err != nil {
		return nil
	}
	return result
}

func (b *BaseRepository) Update(query map[string]interface{}, data interface{}) (count int32, err error) {
	save := b.model.Save(data)
	if save.Error != nil {
		return 0, save.Error
	}
	return int32(save.RowsAffected), nil
}
func (b *BaseRepository) Add(data map[string]interface{}) (int64, error) {
	return 0, nil
}
