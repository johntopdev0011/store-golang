package dtos

type CreateProductRequestDto struct {
	Name        string  `json:"name" validate:"required,gte=0,lte=255"`
	Description string  `json:"description" validate:"required,gte=0,lte=5000"`
	Price       float64 `json:"price" validate:"required,gte=0"`
}
