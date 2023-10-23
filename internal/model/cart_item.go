package model

type CartItem struct {
	CartId    uint
	ProductId uint
	Quantity  uint `gorm:"type:integer;"`
}

func (cartItem *CartItem) Save() error {
	return db.Create(cartItem).Error
}
