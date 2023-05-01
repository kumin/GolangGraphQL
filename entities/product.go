package entities

type Properties struct {
	Id    uint32
	Price float64
	Color string
	Size  string
}

func (p *Properties) TableName() string {
	return "properties"
}

type Product struct {
	Id         uint32
	Name       string
	Sku        string
	Properties *Properties
}

func (p *Product) TableName() string {
	return "products"
}
