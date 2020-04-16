package product

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Repository interface {
	//GetProductById(ProductId int) (*Product, error)
	GetProductById(ProductId string) (*Product, error)
	GetProducts(params *getProductsRequest) ([]*Product, error)
	GetTotalProducts() (int, error)
	InsertProduct(params *getAddProductRequest) (string, error)
	UpdateProduct(params *updateProductRequest) (int, error)
	DeleteProduct(ProductId string) (int, error)
}

type repository struct {
	db *mongo.Database
}

func NewRepository(dbconection *mongo.Database) Repository {
	return &repository{db: dbconection}
}

//func (repo *repository) GetProductById(Product int) (*Product, error) {
func (repo *repository) GetProductById(ProductId string) (*Product, error) {

	//objID, err := primitive.ObjectIDFromHex("5e949091c25c542a557fbb4a")
	objID, err := primitive.ObjectIDFromHex(ProductId)
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

func (repo *repository) InsertProduct(params *getAddProductRequest) (string, error) {
	collection := repo.db.Collection("products")
	// Insert One Document.
	product1 := Product{}
	// An ID for MongoDB.
	product1.ID = primitive.NewObjectID()
	product1.Product_Code = params.Product_Code
	product1.Description = params.Description
	product1.Age, _ = strconv.Atoi(params.Age)
	product1.CreatedAt = time.Now()

	hola, err := collection.InsertOne(context.TODO(), product1)
	fmt.Println(hola)
	if err != nil {
		panic(err)
	}
	return product1.ID.String(), nil

}

func (repo *repository) UpdateProduct(param *updateProductRequest) (int, error) {
	collection := repo.db.Collection("products")
	intAge, _ := strconv.Atoi(param.Age)

	objID, err := primitive.ObjectIDFromHex(param.ID)
	//objID, err := primitive.ObjectIDFromHex("5e94900ab2990b05cc69e6a1")
	if err != nil {
		panic(err)
	}
	resultUpdate, err := collection.UpdateOne(context.TODO(),
		bson.M{"_id": objID},
		bson.M{
			"$set": bson.M{
				"product_code": param.Product_Code,
				"description":  param.Description,
				"age":          intAge,
				//"created_at":   time.Now(),
			},
		},
	)
	return int(resultUpdate.ModifiedCount), nil // output: 1

}

func (repo *repository) DeleteProduct(ProductId string) (int, error) {
	collection := repo.db.Collection("products")

	objID, err := primitive.ObjectIDFromHex(ProductId)
	if err != nil {
		panic(err)
	}
	resultDelete, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		fmt.Println(err)
	}

	return int(resultDelete.DeletedCount), nil
}
