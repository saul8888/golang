package main

import (
	"net/http"
	"time"

	"github.com/GoMongo/database"
	"github.com/GoMongo/product"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DB struct {
	collection *mongo.Collection
}
type getProductRequest struct {
	Limit  int
	Offset int
}

type Product struct {
	ID primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	//ID bson.ObjectId `bson:"_id"`
	//ID           int       `json:"id"`
	Product_Code string    `json:"product_code"`
	Description  string    `json:"description"`
	Age          int       `bson:"age"`
	CreatedAt    time.Time `bson:"created_at"`
}

//var dbconection *mongo.Database

func main() {

	dbconection, _ := database.ConnectDB()
	//fmt.Println(dbconection)
	//collection := dbconection.Collection("products")

	var productRepository = product.NewRepository(dbconection)
	var productService product.Service
	productService = product.NewService(productRepository)
	r := chi.NewRouter()
	r.Mount("/products", product.MakeHttpHandler(productService))
	http.ListenAndServe(":3000", r)

}

//r->routeo q pueden ser GET POST
//r := chi.NewRouter()
//-------------lo q respondo----------lo q recibo
//r.Get("/", AllProducts)
//http.ListenAndServe(":3000", r)

/*
func AllProducts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
	database, err := cliente.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("contet-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

*/
