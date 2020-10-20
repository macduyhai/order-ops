package daos

import (
	"order-ops/dtos"
	"order-ops/models"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type OrderDao interface {
	Create(record *models.Order) error
	Updates(record *models.Order) error
	Search(queries []dtos.SearchQuery) ([]models.Order, error)
	GetByOrderNumber(orderNumber string) (*models.Order, error)
	Delete(orderNumber string) error
	SearchComplate(d time.Time, bra string, typ string, se string, ctr string) ([]models.Order, error)
}

type orderDaoImpl struct {
	db *gorm.DB
}

func NewOrderDao(db *gorm.DB) OrderDao {
	return &orderDaoImpl{db: db}
}

func (dao *orderDaoImpl) Create(record *models.Order) error {
	return dao.db.Create(record).Error
}

func (dao *orderDaoImpl) Updates(record *models.Order) error {
	existedRecord, err := dao.GetByOrderNumber(record.OrderNumber)
	if err != nil {
		return errors.Wrap(err, "get existed record error")
	}

	record.ID = existedRecord.ID
	return dao.db.Model(&existedRecord).Where("id=?", existedRecord.ID).Updates(record).Error
}

//
// func (dao *orderDaoImpl) OrderforDay(timeStart time.Time) (val int64) {

// }

// NumberOrderRequest
func (dao *orderDaoImpl) SearchComplate(d time.Time, bra string, typ string, se string, ctr string) ([]models.Order, error) {
	result := make([]models.Order, 0)
	db := dao.db
	if bra != "" {
		if err := db.Where("branchsell = ? AND status =? AND time_completed=? ", bra, 3, d).Find(&result).Error; err != nil {
			return nil, err
		}
	} else if typ != "" {
		if err := db.Where("typeproduct = ? AND status =?AND time_completed=?  ", typ, 3, d).Find(&result).Error; err != nil {
			return nil, err
		}
	} else if se != "" {
		if err := db.Where("seller = ? AND status =? AND time_completed=? ", se, 3, d).Find(&result).Error; err != nil {
			return nil, err
		}
	} else if ctr != "" {
		if err := db.Where("country = ? AND status =? AND time_completed=? ", ctr, 3, d).Find(&result).Error; err != nil {
			return nil, err
		}
	} else {
		if err := db.Where("status =? AND time_completed=?  ", 3, d).Find(&result).Error; err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (dao *orderDaoImpl) Search(queries []dtos.SearchQuery) ([]models.Order, error) {
	result := make([]models.Order, 0)
	db := dao.db
	for _, query := range queries {
		if query.Key == "status=?" {
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

func (dao *orderDaoImpl) GetByOrderNumber(orderNumber string) (*models.Order, error) {
	var result models.Order
	if err := dao.db.Where("order_number=?", orderNumber).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (dao *orderDaoImpl) Delete(orderNumber string) error {
	var result models.Order
	return dao.db.Where("order_number=?", orderNumber).Delete(&result).Error
}
