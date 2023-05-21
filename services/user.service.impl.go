package services

import (
	"context"
	"errors"

	"example.com/raaho-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	productInfo *mongo.Collection
	ctx         context.Context
}

// custructor
func NewUserService(productInfo *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		productInfo: productInfo,
		ctx:         ctx,
	}
}

//ctx : is used for if we want to perform for certainn perriod of time if taken more than given time process will get cancelled return the appropriate message

func (u *UserServiceImpl) SaveNewProduct(user *models.User) error {
	_, err := u.productInfo.InsertOne(u.ctx, user)
	return err
}

func (u *UserServiceImpl) EditProductDetails(user *models.User) error {
	filter := bson.D{bson.E{Key: "id", Value: user.Id}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "id", Value: user.Id}, bson.E{Key: "sku", Value: user.Sku}, bson.E{Key: "name", Value: user.Name}, bson.E{Key: "description", Value: user.Description}, bson.E{Key: "list_price", Value: user.ListPrice}, bson.E{Key: "sale_price", Value: user.SalePrice}, bson.E{Key: "category", Value: user.Category}, bson.E{Key: "average_product_rating", Value: user.AverageProductRating}, bson.E{Key: "image_url", Value: user.ImageUrl}, bson.E{Key: "brand", Value: user.Brand}, bson.E{Key: "num_reviews", Value: user.NumReviews}}}}

	result, _ := u.productInfo.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("No Matched document Found For Update.")
	}

	return nil
}

func (u *UserServiceImpl) GetProductDetailsByID(id *int) (models.User, error) {
	var user models.User
	query := bson.D{bson.E{Key: "id", Value: id}}
	err := u.productInfo.FindOne(u.ctx, query).Decode(&user)
	return user, err
}

func (u *UserServiceImpl) GetProductsByRating(givenRating *float64) ([]*models.User, error) {
	var users []*models.User
	query := bson.M{
		"average_product_rating": bson.M{
			"$gt": givenRating,
		},
	}
	cursor, err := u.productInfo.Find(u.ctx, query)
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(u.ctx)
	if len(users) == 0 {
		return nil, errors.New("No Match Found")
	}
	return users, nil
}
