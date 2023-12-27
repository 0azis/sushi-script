package domain

type Product struct {
	ID       int    `json:"ID"`
	Name     string `json:"name"`
	Image    string `json:"image"`
	Category string `json:"category"`
	Price    int    `json:"price"`
	Rating   int    `json:"rating"`
}

func (p *Product) Validate() bool {
	if p.Name == "" || p.Image == "" || p.Category == "" || p.Price == 0 {
		return false
	}
	return true
}
