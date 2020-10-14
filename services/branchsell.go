package services

import (
	"fmt"
	"order-ops/daos"
	"order-ops/dtos"
	"order-ops/models"
)

type BranchSellService interface {
	AddBranchSell(request dtos.AddbranchRequest) (*dtos.AddbranchResponse, error)
	// AddLabelsToOrder(request dtos.AddLabelRequest) (*dtos.AddorderResponse, error)
	SearchBranch(queries []dtos.SearchBranchSellQuery) ([]dtos.BranchSell, error)
	// AddShippingTime(request dtos.AddShippingTimeRequest) (*dtos.AddorderResponse, error)
	// MakeCompleted(orderNumber string) (*dtos.AddorderResponse, error)
	// Detete(orderNumber string) error
	// Updates(request dtos.Order) (*dtos.AddorderResponse, error)
}

type branchSellServiceImpl struct {
	dao daos.BranchSellDao
}

func NewBranchSellService(dao daos.BranchSellDao) BranchSellService {
	return &branchSellServiceImpl{
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

func (service *branchSellServiceImpl) mapperDtossToModelBranchSell(input dtos.BranchSell) models.BranchSell {
	return models.BranchSell{
		Name: input.Name,
		Note: input.Note,
	}
}

func (service *branchSellServiceImpl) AddBranchSell(request dtos.AddbranchRequest) (*dtos.AddbranchResponse, error) {
	recordSuccess := make([]string, 0)
	recordFail := make([]string, 0)
	for _, branchsell := range request.BranchSells {
		record := service.mapperDtossToModelBranchSell(branchsell)
		err := service.dao.Create(&record)
		if err != nil {
			recordFail = append(recordFail, branchsell.Name)
		} else {
			recordSuccess = append(recordSuccess, branchsell.Name)
		}
	}

	result := dtos.AddbranchResponse{
		RecordsFailes:  recordFail,
		RecordsSuccess: recordSuccess,
	}

	return &result, nil
}

func (service *branchSellServiceImpl) SearchBranch(queries []dtos.SearchBranchSellQuery) ([]dtos.BranchSell, error) {
	records, _ := service.dao.SearchBranch(queries)
	result := make([]dtos.BranchSell, 0)
	name := ""

	for _, query := range queries {
		if query.Key == "name=?" {
			name = fmt.Sprintf("%v", query.Value)
		}
	}
	for _, record := range records {
		if name != "" {
			if record.Name == name {
				result = append(result, service.mapperDtossToModelBranchSell(record))
			}
		} else {
			result = append(result, service.mapperDtossToModelBranchSell(record))
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
