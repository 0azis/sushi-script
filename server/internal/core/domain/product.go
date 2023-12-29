package domain

type Product struct {
	ID          int                    `json:"ID"`
	Name        string                 `json:"name"`
	Image       string                 `json:"image"`
	Category    int                    `json:"category"`
	Price       int                    `json:"price"`
	Weight      int                    `json:"weight"`
	Description string                 `json:"description"`
	Nutrition   map[string]interface{} `json:"nutrition"`
}

func (p *Product) Validate() bool {
	if p.Name == "" || p.Image == "" || p.Price == 0 {
		return false
	}
	return true
}
