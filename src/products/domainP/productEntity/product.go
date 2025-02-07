package productEntity

type Product struct {
	Id    int32
	Name  string
	Price float32
}

func CreateProduct(_name string, _price float32) *Product {
	return &Product{Name: _name, Price: _price}
}

func (p *Product) SetName(_name string) {
	p.Name = _name
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) SetPrice(_price float32) {
	p.Price = _price
}
func (p *Product) GetPrice() float32 {
	return p.Price
}
