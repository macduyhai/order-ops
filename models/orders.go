package models

import "time"

type Order struct {
	ID                    int64      `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY"`
	OrderNumber           string     `gorm:"column:order_number"`
	CustomerName          string     `gorm:"column:customer_name"`
	Quantity              int32      `gorm:"column:quantity"`
	Phone                 string     `gorm:"column:phone"`
	Address1              string     `gorm:"column:address1"`
	Address2              string     `gorm:"column:address2"`
	City                  string     `gorm:"column:city"`
	State                 string     `gorm:"column:state"`
	PostalCode            string     `gorm:"column:postal_code"`
	Country               string     `gorm:"column:country"`
	TrackingNumber        string     `gorm:"column:tracking_number"`
	URL                   string     `gorm:"column:url"`
	PartnerTrackingNumber string     `gorm:"column:partner_tracking_number"`
	Status                int32      `gorm:"column:status"`
	BranchSell            string     `gorm:"column:branchsell"`
	TypeProduct           string     `gorm:"column:typeproduct"`
	Seller                string     `gorm:"column:seller"`
	Note                  string     `gorm:"column:note"`
	BeginShipping         *time.Time `gorm:"column:begin_shipping"`
	TimeCompleted         *time.Time `gorm:"column:time_completed"`
	CreatedAt             *time.Time `gorm:"column:created_at"`
	UpdatedAt             *time.Time `gorm:"column:updated_at"`
	DeletedAt             *time.Time `gorm:"column:deleted_at"`
}

func (Order) TableName() string {
	return "orders"
}

type Item struct {
	ID          int64      `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY"`
	OrderNumber string     `gorm:"column:order_number"`
	TypeProduct string     `gorm:"column:typeproduct"`
	Quantity    int32      `gorm:"column:quantity"`
	Note        string     `gorm:"column:note"`
	CreatedAt   *time.Time `gorm:"column:created_at"`
}
