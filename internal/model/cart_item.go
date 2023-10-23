package model

type CartItem struct {
	ID        uint `gorm:"primaryKey"`
	CartId    uint
	ProductId uint
	Quantity  uint `gorm:"type:integer;"`
}

func (cartItem *CartItem) Save() error {
	return db.Create(cartItem).Error
}

func (cartItem *CartItem) Update() error {
	return db.Save(cartItem).Error
}
