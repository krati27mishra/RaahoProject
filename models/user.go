package models

type User struct {
	Id                   int     `json:"id" bson:"id"`
	Sku                  string  `json:"sku" bson:"sku"`
	Name                 string  `json:"name" bson:"name"`
	Description          string  `json:"description" bson:"description"`
	ListPrice            float64 `json:"list_price" bson:"list_price"`
	SalePrice            float64 `json:"sale_price" bson:"sale_price"`
	Category             string  `json:"category" bson:"category"`
	AverageProductRating float64 `json:"average_product_rating" bson:"average_product_rating"`
	ImageUrl             string  `json:"image_url" bson:"image_url"`
	Brand                string  `json:"brand" bson:"brand"`
	NumReviews           int     `json:"num_reviews" bson:"num_reviews"`
}
