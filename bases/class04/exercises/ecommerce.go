package exercises

/* START PRODUCT */

const (
	small  = "small"
	medium = "medium"
	large  = "large"
)

type Product struct {
	Name  string
	Size  string
	Price float32
}

func NewProduct(name, size string, price float32) Product {
	return Product{name, size, price}
}

type Producter interface {
	CalcTotalWithTaxes() float32
}

func (p Product) CalcTotalWithTaxes() float32 {
	switch p.Size {
	case medium:
		return (p.Price * 0.03) + p.Price
	case large:
		return ((p.Price * 0.06) + p.Price) + 2500
	default:
		return p.Price
	}
}

/* END PRODUCT */

/* START CART */

type Carter interface {
	Add(Product)
	Total()
}

type Cart struct {
	Products []Product
}

func NewCart() Cart {
	return Cart{}
}

func (c *Cart) Add(p Product) {
	c.Products = append(c.Products, p)
}

func (c Cart) Total() (total float32) {
	for _, p := range c.Products {
		total += p.CalcTotalWithTaxes()
	}
	return
}

/* END CART */
