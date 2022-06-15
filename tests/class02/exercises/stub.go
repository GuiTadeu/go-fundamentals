package exercises

type StubSearchEngine struct {}

func NewStubEngine(s StubSearchEngine) *StubSearchEngine {
	return &StubSearchEngine{}
}

type StubProduct struct {
	name string
	price float32
}

func (d StubSearchEngine) StubGetAll() []StubProduct {
	return []StubProduct{
		{"Calça para uma jovem", 350.00},
		{"Disco do O Terno", 66.00},
	}
}