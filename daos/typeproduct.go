package daos

import (
	"order-ops/models"

	"github.com/jinzhu/gorm"
)

type TypeProductDao interface {
	Create(record *models.TypeProduct) error
	// Updates(record *models.Order) error
	// Search(queries []dtos.SearchQuery) ([]models.Order, error)
	// GetByOrderNumber(orderNumber string) (*models.Order, error)
	// Delete(orderNumber string) error
}

type typeProductDaoImpl struct {
	db *gorm.DB
}

func NewTypeProductDao(db *gorm.DB) TypeProductDao {
	return &typeProductDaoImpl{db: db}
}

func (dao *typeProductDaoImpl) Create(record *models.TypeProduct) error {
	return dao.db.Create(record).Error
}

// func (dao *orderDaoImpl) Updates(record *models.Order) error {
// 	existedRecord, err := dao.GetByOrderNumber(record.OrderNumber)
// 	if err != nil {
// 		return errors.Wrap(err, "get existed record error")
// 	}

// 	record.ID = existedRecord.ID
// 	return dao.db.Model(&existedRecord).Where("id=?", existedRecord.ID).Updates(record).Error
// }

// func (dao *orderDaoImpl) Search(queries []dtos.SearchQuery) ([]models.Order, error) {
// 	result := make([]models.Order, 0)
// 	db := dao.db
// 	for _, query := range queries {
// 		if query.Key == "status=?" {
// 			continue
// 		}

// 		if query.Value != nil {
// 			db = db.Where(query.Key, query.Value)
// 		} else {
// 			db = db.Where(query.Key)
// 		}
// 	}

// 	if err := db.Find(&result).Error; err != nil {
// 		return nil, nil
// 	}

// 	return result, nil
// }

// func (dao *orderDaoImpl) GetByOrderNumber(orderNumber string) (*models.Order, error) {
// 	var result models.Order
// 	if err := dao.db.Where("order_number=?", orderNumber).First(&result).Error; err != nil {
// 		return nil, err
// 	}

// 	return &result, nil
// }

// func (dao *orderDaoImpl) Delete(orderNumber string) error {
// 	var result models.Order
// 	return dao.db.Where("order_number=?", orderNumber).Delete(&result).Error
// }
