package models

type Product struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	CategoryID   uint    `json:"category_id"`
	CostPrice    float64 `json:"cost_price"`
	SellingPrice float64 `json:"selling_price"`
}

type ProductRevenueResponse struct {
	Product string  `json:"product"`
	Revenue float64 `json:"revenue"`
}

type ProductCreateRequest struct {
	Start string `json:"start" Validate:"required"`
	End   string `json:"end" Validate:"required"`
}
