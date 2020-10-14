package daos

import (
	"order-ops/dtos"
	"order-ops/models"

	"github.com/jinzhu/gorm"
)

type BranchSellDao interface {
	Create(record *models.BranchSell) error
	// Updates(record *models.Order) error
	SearchBranch(queries []dtos.SearchBranchSellQuery) ([]dtos.BranchSell, error)
	// GetByOrderNumber(orderNumber string) (*models.Order, error)
	// Delete(orderNumber string) error
}

type branchSellDaoImpl struct {
	db *gorm.DB
}

func NewBranchSellDao(db *gorm.DB) BranchSellDao {
	return &branchSellDaoImpl{db: db}
}

func (dao *branchSellDaoImpl) Create(record *models.BranchSell) error {
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

func (dao *branchSellDaoImpl) SearchBranch(queries []dtos.SearchBranchSellQuery) ([]dtos.BranchSell, error) {
	result := make([]dtos.BranchSell, 0)
	db := dao.db
	for _, query := range queries {
		if query.Key == "name=?" {
			continue
		}

		if query.Value != nil {
			db = db.Where(query.Key, query.Value)
		} else {
			db = db.Where(query.Key)
		}
	}

	if err := db.Find(&result).Error; err != nil {
		return nil, nil
	}

	return result, nil
}

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
