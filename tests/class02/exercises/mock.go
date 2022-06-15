package exercises

var products = []MockProduct{
	{1, "Before Update", 350.00},
	{2, "Disco do O Terno", 66.00},
}

type MockSearchEngine struct {
	GetAllWasCalled bool
}

func NewMockEngine(mock MockSearchEngine) MockSearchEngine {
	return mock
}

type MockProduct struct {
	id uint64
	name string
	price float32
}

func (m *MockSearchEngine) UpdateName(id uint64, newName string) {

	for index, product := range m.MockGetAll() {
		if(product.id == id) {
			p := &products[index]
			p.name = newName
		}
	}
}

func (m *MockSearchEngine) MockGetAll() []MockProduct {
	m.GetAllWasCalled = true
	return products
}