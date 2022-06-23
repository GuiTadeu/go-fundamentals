package class04

type Transaction struct {
	Id        uint64
	Code      string
	Currency  string
	Value     float64
	Issuer    string
	Receiver  string
	CreatedAt string
}