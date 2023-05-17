package entities

type Properties struct {
	ID        uint32  `json:"id,omitempty" gorm:"primaryKey"`
	Price     float64 `json:"price,omitempty"`
	Color     string  `json:"color,omitempty"`
	Size      string  `json:"size,omitempty"`
	ProductId uint32  `json:"product_id,omitempty"`
}

func (p *Properties) TableName() string {
	return "properties"
}

type Product struct {
	ID         uint32      `json:"id,omitempty" gorm:"primaryKey"`
	Name       string      `json:"name,omitempty"`
	Sku        string      `json:"sku,omitempty"`
	Properties *Properties `json:"properties,omitempty"`
}

func (p *Product) TableName() string {
	return "products"
}
