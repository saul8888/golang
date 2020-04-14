package product

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	//ID                bson.ObjectId `bson:"_id"`
	ID primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	//ID           int       `json:"id"`
	Product_Code string    `json:"product_code"`
	Description  string    `json:"description"`
	Age          int       `bson:"age"`
	CreatedAt    time.Time `bson:"created_at"`
}

type ProductList struct {
	Data         []*Product `json:"data"`
	TotalRecords int        `json:"totalRecords"`
}
