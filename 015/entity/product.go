package entity

type Product struct {
	ID    int
	Name  string
	Price int
	Stock int
}

func (p Product) StockStatus() string {
	var status string
	if p.Stock < 3 {
		status = "Grab it fast! Stock almost empty"
	} else if p.Stock < 10 {
		status = "Limited stock!"
	}

	return status
}
