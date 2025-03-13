package dtos

// CreatePersonRequest defines the POST /people request body
type CreatePersonRequest struct {
	Name string `json:"name" validate:"required,min=5,capitalized"`
	Age  int    `json:"age,omitempty" validate:"gte=0"` // Optional field
}

// PersonResponse defines the GET /people response item
type PersonResponse struct {
	Name string `json:"name"`
}
