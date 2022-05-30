package exercises

type Customer struct {
	Name     string
	Surname  string
	Email    string
	Products []Product
}

type Product struct {
	name     string
	price    float32
	quantity uint16
}

func Build(name string, price float32, quantity uint16) Product {
	return Product{name, price, quantity}
}

func AddProduct(c *Customer, p Product) {
	c.Products = append(c.Products, p)
}

func ClearProducts(c *Customer) {
	c.Products = []Product{}
}