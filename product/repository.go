package product

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Repository interface {
	GetProductById(ProductId int) (*Product, error)
	GetProducts(params *getProductsRequest) ([]*Product, error)
	GetTotalProducts() (int, error)
}

type repository struct {
	//db     *sql.DB
	db *mongo.Database
}

//func NewRepository(dataconection *sql.DB) Repository {
//	return &repository{db: dataconection}
//}
func NewRepository(dbconection *mongo.Database) Repository {
	return &repository{db: dbconection}
}

//func (repo *repository) GetProductById(Product int) (*Product, error) {
//
//}
func (repo *repository) GetProductById(ProductId int) (*Product, error) {

	objID, err := primitive.ObjectIDFromHex("5e94907372796da9a4751d76")
	if err != nil {
		panic(err)
	}

	collection := repo.db.Collection("products")
	row, err := collection.Find(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		fmt.Println(err)
	}

	product := &Product{}
	product2 := []Product{}

	for row.Next(context.TODO()) {
		row.Decode(&product)
		product2 = append(product2, *product)
	}
	//for _, eldato := range product2 {
	//	fmt.Println(eldato)
	//}

	return product, err

}

func (repo *repository) GetProducts(params *getProductsRequest) ([]*Product, error) {
	collection := repo.db.Collection("products")
	options := options.Find()
	options.SetSkip(int64(params.Offset))
	options.SetLimit(int64(params.Limit))
	row, err := collection.Find(context.TODO(), bson.M{}, options)
	if err != nil {
		fmt.Println(err)
	}

	//product := &Product{}
	var product2 []*Product

	for row.Next(context.TODO()) {
		product := &Product{}
		row.Decode(&product)
		product2 = append(product2, product)
	}
	//for _, eldato := range product2 {
	//	fmt.Println(eldato)
	//}

	return product2, err

}

func (repo *repository) GetTotalProducts() (int, error) {
	collection := repo.db.Collection("products")
	total, err := collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}
	return int(total), nil
}
