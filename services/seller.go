package services

import (
	"fmt"
	"order-ops/daos"
	"order-ops/dtos"
	"order-ops/models"
)

type SellerService interface {
	AddSeller(request dtos.AddsellerRequest) (*dtos.AddsellerResponse, error)
	// AddLabelsToOrder(request dtos.AddLabelRequest) (*dtos.AddorderResponse, error)
	SearchSeller(queries []dtos.SearchSellerQuery) ([]dtos.Seller, error)
	// AddShippingTime(request dtos.AddShippingTimeRequest) (*dtos.AddorderResponse, error)
	// MakeCompleted(orderNumber string) (*dtos.AddorderResponse, error)
	// Detete(orderNumber string) error
	// Updates(request dtos.Order) (*dtos.AddorderResponse, error)
}

type sellerServiceImpl struct {
	dao daos.SellerDao
}

func NewSellerService(dao daos.SellerDao) SellerService {
	return &sellerServiceImpl{
		dao: dao,
	}
}

func (service *sellerServiceImpl) mapperDtossToModelSeller(input dtos.Seller) models.Seller {
	return models.Seller{
		Name: input.Name,
		Note: input.Note,
	}
}

func (service *sellerServiceImpl) mapperModelToDtossSeller(input models.Seller) dtos.Seller {
	return dtos.Seller{
		Name: input.Name,
		Note: input.Note,
	}
}

func (service *sellerServiceImpl) AddSeller(request dtos.AddsellerRequest) (*dtos.AddsellerResponse, error) {
	recordSuccess := make([]string, 0)
	recordFail := make([]string, 0)
	for _, seller := range request.Sellers {
		record := service.mapperDtossToModelSeller(seller)
		err := service.dao.Create(&record)
		if err != nil {
			recordFail = append(recordFail, seller.Name)
		} else {
			recordSuccess = append(recordSuccess, seller.Name)
		}
	}

	result := dtos.AddsellerResponse{
		RecordsFailes:  recordFail,
		RecordsSuccess: recordSuccess,
	}

	return &result, nil
}

func (service *sellerServiceImpl) SearchSeller(queries []dtos.SearchSellerQuery) ([]dtos.Seller, error) {
	records, _ := service.dao.SearchSeller(queries)
	result := make([]dtos.Seller, 0)
	name := ""

	for _, query := range queries {
		if query.Key == "name=?" {
			name = fmt.Sprintf("%v", query.Value)
		}
	}
	for _, record := range records {
		if name != "" {
			if record.Name == name {
				result = append(result, service.mapperModelToDtossSeller(record))
			}
		} else {
			result = append(result, service.mapperModelToDtossSeller(record))
		}
	}

	return result, nil
}

// func (service *orderServiceImpl) AddShippingTime(request dtos.AddShippingTimeRequest) (*dtos.AddorderResponse, error) {
// 	record := models.Order{
// 		OrderNumber:   request.OrderNumber,
// 		BeginShipping: request.BeginShippingReal,
// 		TimeCompleted: request.TimeCompletedReal,
// 	}
// 	err := service.dao.Updates(&record)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &dtos.AddorderResponse{
// 		ID: record.ID,
// 	}, nil
// }

// func (service *orderServiceImpl) MakeCompleted(orderNumber string) (*dtos.AddorderResponse, error) {
// 	record := models.Order{
// 		OrderNumber: orderNumber,
// 		Status:      completedStatus,
// 	}
// 	err := service.dao.Updates(&record)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &dtos.AddorderResponse{
// 		ID: record.ID,
// 	}, nil
// }

// func (service *orderServiceImpl) Detete(orderNumber string) error {
// 	return service.dao.Delete(orderNumber)
// }

// func (service *orderServiceImpl) Updates(request dtos.Order) (*dtos.AddorderResponse, error) {
// 	record := service.mapperDtossToModelOrder(request)
// 	err := service.dao.Updates(&record)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "update record error")
// 	}

// 	return &dtos.AddorderResponse{ID: record.ID}, nil
// }
