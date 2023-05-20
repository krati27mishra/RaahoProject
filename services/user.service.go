package services

import "example.com/raaho-api/models"

type UserService interface {
	SaveNewProduct(*models.User) error
	EditProductDetails(*models.User) error
	GetProductDetailsByID(*int) (*models.User, error)
	GetProductsByRating(*int) ([]*models.User, error)
}
