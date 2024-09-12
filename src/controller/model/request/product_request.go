package request

type ProductRequest struct {
	Id 			string `json:"id" binding:"required"`
	Name		string `json:"name" binding:"required,min=4,max=50"`
	Price 		float64 `json:"price" binding:"required"`
	Quantity	int `json:"quantity" binding:"required,numeric,min=1,max=100"`
}