package repository

import (
	"github.com/jinzhu/gorm"
	"log"
)

type Reader interface {
	Find(id int32) (interface{}, error)
	Select(query map[string]interface{}) map[string]interface{}
}

type Writer interface {
	Update(query map[string]interface{}, data interface{}) (count int32, err error)
	Add(data map[string]interface{}) (int64, error)
}

type Repository interface {
	Reader
	Writer
}
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
	log.Println("开始查询")
	result := map[string]interface{}{}

	err := b.model.Where(query).First(&result).Error
	log.Println(err)
	if err != nil {
		return nil
	}
	//toMap, err := apgs.ToMap(result, "json")
	//log.Println(err)
	//if err != nil {
	//	return nil
	//}
	log.Println("结束查询")
	log.Println(&result)
	log.Println(result)
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
