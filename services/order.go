package services

import (
	"fmt"
	"log"
	"order-ops/daos"
	"order-ops/dtos"
	"order-ops/models"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type OrderService interface {
	AddOrder(request dtos.AddOrderRequest) (*dtos.AddorderResponse, error)
	AddLabelsToOrder(request dtos.AddLabelRequest) (*dtos.AddorderResponse, error)
	AddLabelsToItems(request dtos.AddLabelRequest) (*dtos.AddorderResponse, error)
	Search(queries []dtos.SearchQuery) ([]dtos.FullOrderInformation, error)
	SearchItems(queries []dtos.SearchItemsQuery) ([]dtos.Item, error)
	AddShippingTime(request dtos.AddShippingTimeRequest) (*dtos.AddorderResponse, error)
	MakeCompleted(orderNumber string) (*dtos.AddorderResponse, error)
	MakeDelay(orderNumber string) (*dtos.AddorderResponse, error)
	Detete(orderNumber string) error
	Updates(request dtos.Order) (*dtos.AddorderResponse, error)
	// NumberOrders(request dtos.NumberOrderRequest) (*dtos.NumberOrderResponse, error)
}

type orderServiceImpl struct {
	dao daos.OrderDao
}

func NewOrderService(dao daos.OrderDao) OrderService {
	return &orderServiceImpl{
		dao: dao,
	}
}

const CommonTimeFormat = "2006-01-02 15:04:05"
const (
	waitingStatus   = 0
	shippingStatus  = 2
	completedStatus = 3
	delayStatus     = 1

	// waitingStatus   = 0
	// shippingStatus  = 1
	// completedStatus = 2
	// delayStatus     = 3

	// processingStatus = 0
	// shippingStatus   = 1
	// holdOnStatus     = 2
	// completedStatus  = 3
)

func (service *orderServiceImpl) Filtercharacter(input string) string {
	new := ""
	if input != "" {
		new = strings.ToLower(input)
		new = strings.Replace(new, " ", "", -1)
	}
	return new
}

func (service *orderServiceImpl) mapperDtossToModelOrder(input dtos.Order) models.Order {
	// input.Country = service.Filtercharacter(input.Country)
	input.BranchSell = service.Filtercharacter(input.BranchSell)
	input.TypeProduct = service.Filtercharacter(input.TypeProduct)
	input.Seller = service.Filtercharacter(input.Seller)
	t_n := time.Now().Add(+7 * time.Hour)

	return models.Order{
		OrderNumber:  input.OrderNumber,
		CustomerName: input.Name,
		Quantity:     input.Quantity,
		Phone:        input.Phone,
		Address1:     input.Address1,
		Address2:     input.Address2,
		City:         input.City,
		State:        input.State,
		PostalCode:   input.PostalCode,
		Country:      input.Country,
		BranchSell:   input.BranchSell,
		TypeProduct:  input.TypeProduct,
		Seller:       input.Seller,
		Note:         input.Note,
		CreatedAt:    &t_n,
		PrintStatus:  input.PrintStatus,
	}
}

func (service *orderServiceImpl) mapperDtossToModelOrderAddLable(input dtos.AddLabelRequest) models.Order {
	return models.Order{
		OrderNumber:    input.OrderNumber,
		TrackingNumber: input.LableDetails.TrackingNumber,
		URL:            input.LableDetails.URL,
		PartnerTrackingNumber: input.LableDetails.PartnerTrackingNumber,
	}
}

func (service *orderServiceImpl) AddOrder(request dtos.AddOrderRequest) (*dtos.AddorderResponse, error) {
	recordSuccess := make([]string, 0)
	recordFail := make([]string, 0)
	for _, order := range request.Orders {
		if order.OrderNumber != "" {
			order.Country = strings.ToUpper(order.Country)
			record := service.mapperDtossToModelOrder(order)
			err := service.dao.Create(&record)
			if err != nil {
				recordFail = append(recordFail, order.OrderNumber)
			} else {
				recordSuccess = append(recordSuccess, order.OrderNumber)
			}
		} else {
			recordFail = append(recordFail, order.OrderNumber)
		}
	}

	result := dtos.AddorderResponse{
		RecordsFailes:  recordFail,
		RecordsSuccess: recordSuccess,
	}

	return &result, nil
}

func (service *orderServiceImpl) AddLabelsToOrder(request dtos.AddLabelRequest) (*dtos.AddorderResponse, error) {
	record := service.mapperDtossToModelOrderAddLable(request)
	err := service.dao.Updates(&record)
	if err != nil {
		return nil, err
	}

	return &dtos.AddorderResponse{
		ID: record.ID,
	}, nil
}
func (service *orderServiceImpl) mapperDtossToModelItemAddLable(input dtos.Item) models.Item {
	t_n := time.Now().Add(+7 * time.Hour)
	return models.Item{
		OrderNumber:      input.OrderNumber,
		SkuNumber:        input.SkuNumber,
		PackagedQuantity: input.PackagedQuantity,
		ItemDescription:  input.ItemDescription,
		CreatedAt:        &t_n,
	}
}

//  mapperDtossToModelItemAddLable(input dtos.Item) models.Item
func (service *orderServiceImpl) AddLabelsToItems(request dtos.AddLabelRequest) (*dtos.AddorderResponse, error) {
	res := dtos.AddorderResponse{}
	log.Println(request)
	for _, item := range request.Items {
		if item.SkuNumber != "" {
			if item.OrderNumber == "" {
				item.OrderNumber = request.OrderNumber
			}
			record := service.mapperDtossToModelItemAddLable(item)
			err := service.dao.Create_Item(&record)
			if err != nil {
				return nil, err
			}
		}

	}
	return &res, nil
}

func (service *orderServiceImpl) mapperModelsToOrderFullInfor(input models.Order) dtos.FullOrderInformation {
	begin := input.BeginShipping.Format(CommonTimeFormat)
	end := input.TimeCompleted.Format(CommonTimeFormat)
	if begin == end {
		begin = ""
		end = ""
	}

	return dtos.FullOrderInformation{
		dtos.Order{
			OrderNumber: input.OrderNumber,
			Name:        input.CustomerName,
			Quantity:    input.Quantity,
			Phone:       input.Phone,
			Address1:    input.Address1,
			Address2:    input.Address2,
			City:        input.City,
			State:       input.State,
			PostalCode:  input.PostalCode,
			Country:     input.Country,
			BranchSell:  input.BranchSell,
			TypeProduct: input.TypeProduct,
			Seller:      input.Seller,
			Note:        input.Note,
			CreatedAt:   input.CreatedAt,
			PrintStatus: input.PrintStatus,
		},
		dtos.ShippingInfor{
			Status:        input.Status,
			BeginShipping: begin,
			TimeCompleted: end,
		},
		dtos.LableDetails{
			TrackingNumber: input.TrackingNumber,
			URL:            input.URL,
			PartnerTrackingNumber: input.PartnerTrackingNumber,
		},
	}
}

func (service *orderServiceImpl) updateRecordState(input *models.Order) {
	if input.BeginShipping.Equal(*input.TimeCompleted) || input.Status == completedStatus {
		return
	}

	now := time.Now()
	now = now.Add(+7 * time.Hour)
	log.Println(now)
	// if now.Equal(*input.BeginShipping) {
	// 	input.Status = shippingStatus
	// 	return
	// }
	if now.After(*input.BeginShipping) && now.Before(*input.TimeCompleted) {
		input.Status = shippingStatus
		return
	}

	if now.After(*input.BeginShipping) {
		input.Status = completedStatus
		return
	}
	log.Println(input.Status)
}

func (service *orderServiceImpl) mapperModelToDtossItem(input models.Item) dtos.Item {
	t_n := time.Now().Add(+7 * time.Hour)
	return dtos.Item{
		OrderNumber:      input.OrderNumber,
		SkuNumber:        input.SkuNumber,
		PackagedQuantity: input.PackagedQuantity,
		ItemDescription:  input.ItemDescription,
		CreatedAt:        &t_n,
	}
}
func (service *orderServiceImpl) SearchItems(queries []dtos.SearchItemsQuery) ([]dtos.Item, error) {
	records, _ := service.dao.SearchItems(queries)
	result := make([]dtos.Item, 0)

	for _, record := range records {
		if record.OrderNumber != "" {
			result = append(result, service.mapperModelToDtossItem(record))
		}
	}

	return result, nil
}

//
func (service *orderServiceImpl) Search(queries []dtos.SearchQuery) ([]dtos.FullOrderInformation, error) {
	records, _ := service.dao.Search(queries)
	result := make([]dtos.FullOrderInformation, 0)
	status := -1

	for _, query := range queries {
		if query.Key == "status=?" {
			statusint, _ := strconv.Atoi(fmt.Sprintf("%v", query.Value))
			status = statusint
		}
	}

	for _, record := range records {
		if int(record.Status) != delayStatus {
			service.updateRecordState(&record)
		}
		if status != -1 {
			if int(record.Status) == status {
				result = append(result, service.mapperModelsToOrderFullInfor(record))
			}
		} else {
			result = append(result, service.mapperModelsToOrderFullInfor(record))
		}
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].Status < result[j].Status
	})

	return result, nil
}

func (service *orderServiceImpl) AddShippingTime(request dtos.AddShippingTimeRequest) (*dtos.AddorderResponse, error) {
	record := models.Order{
		OrderNumber:   request.OrderNumber,
		BeginShipping: request.BeginShippingReal,
		TimeCompleted: request.TimeCompletedReal,
	}
	err := service.dao.Updates(&record)
	if err != nil {
		return nil, err
	}

	return &dtos.AddorderResponse{
		ID: record.ID,
	}, nil
}

func (service *orderServiceImpl) MakeDelay(orderNumber string) (*dtos.AddorderResponse, error) {
	record := models.Order{
		OrderNumber: orderNumber,
		Status:      delayStatus,
	}
	err := service.dao.Updates(&record)
	if err != nil {
		return nil, err
	}

	return &dtos.AddorderResponse{
		ID: record.ID,
	}, nil
}

func (service *orderServiceImpl) MakeCompleted(orderNumber string) (*dtos.AddorderResponse, error) {
	tg := time.Now()
	tg = tg.Add(+7 * time.Hour)

	record := models.Order{
		OrderNumber:   orderNumber,
		Status:        completedStatus,
		TimeCompleted: &tg,
	}
	err := service.dao.Updates(&record)
	if err != nil {
		return nil, err
	}

	return &dtos.AddorderResponse{
		ID: record.ID,
	}, nil
}

func (service *orderServiceImpl) Detete(orderNumber string) error {
	return service.dao.Delete(orderNumber)
}

func (service *orderServiceImpl) Updates(request dtos.Order) (*dtos.AddorderResponse, error) {
	record := service.mapperDtossToModelOrder(request)
	err := service.dao.Updates(&record)
	if err != nil {
		return nil, errors.Wrap(err, "update record error")
	}

	return &dtos.AddorderResponse{ID: record.ID}, nil
}
