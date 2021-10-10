package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/shunyaYoshimra/checkCORS/database"
)

type MaterialRepository struct {
	Conn *gorm.DB
}

func NewMaterialRepository() MaterialRepository {
	return MaterialRepository{Conn: database.GetDB().Table("materials")}
}

func (mr *MaterialRepository) RetrieveMaterials() (materials []database.Material) {
	mr.Conn.Order("created_at desc").Find(&materials)
	return
}

func (mr *MaterialRepository) RetrieveMaterialsByKeyword(keyword string) (materials []database.Material) {
	mr.Conn.Where("tags LIKE ?", "%"+keyword+"%").Order("created_at desc").Find(&materials)
	return
}

func (mr *MaterialRepository) Create(m *database.Material) (err error) {
	err = mr.Conn.Create(m).Error
	return
}

func (mr *MaterialRepository) Delete(m database.Material) (err error) {
	err = mr.Conn.Delete(&m).Error
	return
}