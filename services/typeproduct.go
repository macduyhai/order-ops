package services

import (
	"fmt"
	"order-ops/daos"
	"order-ops/dtos"
	"order-ops/models"
)

type TypeProductService interface {
	AddTypeProduct(request dtos.AddtypeRequest) (*dtos.AddtypeResponse, error)
	// AddLabelsToOrder(request dtos.AddLabelRequest) (*dtos.AddorderResponse, error)
	SearchType(queries []dtos.SearchTypeProductQuery) ([]dtos.TypeProduct, error)
	// AddShippingTime(request dtos.AddShippingTimeRequest) (*dtos.AddorderResponse, error)
	// MakeCompleted(orderNumber string) (*dtos.AddorderResponse, error)
	// Detete(orderNumber string) error
	// Updates(request dtos.Order) (*dtos.AddorderResponse, error)
}

type typeProductServiceImpl struct {
	dao daos.TypeProductDao
}

func NewTypeProductService(dao daos.TypeProductDao) TypeProductService {
	return &typeProductServiceImpl{
		dao: dao,
	}
}

// const CommonTimeFormat = "2006-01-02 15:04:05"
// const (
// 	processingStatus = 0
// 	shippingStatus   = 1
// 	holdOnStatus     = 2
// 	completedStatus  = 3
// )

func (service *typeProductServiceImpl) mapperDtossToModelTypeProduct(input dtos.TypeProduct) models.TypeProduct {
	return models.TypeProduct{
		Name:   input.Name,
		Width:  input.Width,
		Height: input.Height,
		Weight: input.Weight,
		Length: input.Length,
		Note:   input.Note,
	}
}
func (service *typeProductServiceImpl) mapperModelToDtossTypeProduct(input models.TypeProduct) dtos.TypeProduct {
	return dtos.TypeProduct{
		Name:   input.Name,
		Width:  input.Width,
		Height: input.Height,
		Weight: input.Weight,
		Length: input.Length,
		Note:   input.Note,
	}
}

// func (service *orderServiceImpl) mapperDtossToModelOrderAddLable(input dtos.AddLabelRequest) models.Order {
// 	return models.Order{
// 		OrderNumber:    input.OrderNumber,
// 		TrackingNumber: input.LableDetails.TrackingNumber,
// 		URL:            input.LableDetails.URL,
// 		PartnerTrackingNumber: input.LableDetails.PartnerTrackingNumber,
// 	}
// }

func (service *typeProductServiceImpl) AddTypeProduct(request dtos.AddtypeRequest) (*dtos.AddtypeResponse, error) {
	recordSuccess := make([]string, 0)
	recordFail := make([]string, 0)
	for _, typeproduct := range request.TypeProducts {
		record := service.mapperDtossToModelTypeProduct(typeproduct)
		err := service.dao.Create(&record)
		if err != nil {
			recordFail = append(recordFail, typeproduct.Name)
		} else {
			recordSuccess = append(recordSuccess, typeproduct.Name)
		}
	}

	result := dtos.AddtypeResponse{
		RecordsFailes:  recordFail,
		RecordsSuccess: recordSuccess,
	}

	return &result, nil
}

func (service *typeProductServiceImpl) SearchType(queries []dtos.SearchTypeProductQuery) ([]dtos.TypeProduct, error) {
	records, _ := service.dao.SearchType(queries)
	result := make([]dtos.TypeProduct, 0)
	name := ""

	for _, query := range queries {
		if query.Key == "name=?" {
			name = fmt.Sprintf("%v", query.Value)
		}
	}
	for _, record := range records {
		if name != "" {
			if record.Name == name {
				result = append(result, service.mapperModelToDtossTypeProduct(record))
			}
		} else {
			result = append(result, service.mapperModelToDtossTypeProduct(record))
		}
	}

	return result, nil
}
func (service *typeProductServiceImpl) Detete(Typename string) error {
	return service.dao.Delete(Typename)
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

// func (service *orderServiceImpl) Updates(request dtos.Order) (*dtos.AddorderResponse, error) {
// 	record := service.mapperDtossToModelOrder(request)
// 	err := service.dao.Updates(&record)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "update record error")
// 	}

// 	return &dtos.AddorderResponse{ID: record.ID}, nil
// }
