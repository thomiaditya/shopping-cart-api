package model

import "gorm.io/gorm"

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusCompleted OrderStatus = "completed"
)

type Order struct {
	gorm.Model
	CustomerId uint         `gorm:"not null"`
	CartId     uint         `gorm:"not null"`
	PaymentId  uint         `gorm:"not null"`
	Status     *OrderStatus `gorm:"type:varchar(100);default:pending"`

	Customer *Customer `gorm:"foreignKey:CustomerId"`
	Cart     *Cart     `gorm:"foreignKey:CartId"`
	Payment  *Payment  `gorm:"foreignKey:PaymentId"`
}

func (Order) TableName() string {
	return "orders"
}

func CreateNewOrder(order *Order) (*Order, error) {
	err := order.Save()
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (order *Order) Save() error {
	return db.Create(order).Error
}

func (order *Order) Update() error {
	return db.Save(order).Error
}
