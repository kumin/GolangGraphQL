package entities

type Properties struct {
	Id        uint32 `gorm:"primaryKey"`
	Price     float64
	Color     string
	Size      string
	ProductId uint32
}

func (p *Properties) TableName() string {
	return "properties"
}

type Product struct {
	Id         uint32 `gorm:"primaryKey"`
	Name       string
	Sku        string
	Properties *Properties
}

func (p *Product) TableName() string {
	return "products"
}
