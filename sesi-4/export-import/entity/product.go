package entity

type Product struct {
	Name  string
	Price float64
	owner string
}

func (p Product) GetOwnerName() string {
	p.owner = "John Doe"

	return p.owner
}
