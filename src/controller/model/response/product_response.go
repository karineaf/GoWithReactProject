package response

type ProductResponse struct {
	Id 			string `json:"id"`
	Name		string `json:"name"`
	Price 		float64 `json:"price"`
	Quantity	int `json:"quantity"`
}