package product

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getProductByIDRequest struct {
	//ProductID int
	ProductID string
}

type getProductsRequest struct {
	Limit  int
	Offset int
}

type getAddProductRequest struct {
	//ID           primitive.ObjectID
	//ID           string
	Product_Code string
	Description  string
	Age          string
	//CreatedAt    string
	//CreatedAt    time.Time
}

type updateProductRequest struct {
	ID           string
	Product_Code string
	Description  string
	Age          string
}

type deleteProductRequest struct {
	ProductID string
}

func makeGetProductByIdEndPoint(s Service) endpoint.Endpoint {
	getProductByIdEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductByIDRequest)
		product, err := s.GetProductById(&req)
		if err != nil {
			panic(err)
		}
		return product, nil
	}

	return getProductByIdEndpoint
}

func makeGetProductsEndPoint(s Service) endpoint.Endpoint {
	getProductsEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductsRequest)
		result, err := s.GetProducts(&req)
		if err != nil {
			panic(err)
		}
		return result, nil
	}
	return getProductsEndPoint
}

func makeAddProductEndpoint(s Service) endpoint.Endpoint {
	addProductEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getAddProductRequest)
		result, err := s.InsertProduct(&req)
		if err != nil {
			panic(err)
		}
		return result, nil
	}
	return addProductEndpoint
}

func makeUpdateProductEndpoint(s Service) endpoint.Endpoint {
	updateProductEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateProductRequest)
		result, err := s.UpdateProduct(&req)
		if err != nil {
			panic(err)
		}
		return result, nil
	}
	return updateProductEndpoint
}

func makeDeleteProductEndpoint(s Service) endpoint.Endpoint {
	deleteProductEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteProductRequest)
		result, err := s.DeleteProduct(&req)
		if err != nil {
			panic(err)
		}
		return result, nil
	}
	return deleteProductEndpoint
}
