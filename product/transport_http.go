package product

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()

	getProductByHandler := kithttp.NewServer(
		makeGetProductByIdEndPoint(s),
		getProductByIdRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/{id}", getProductByHandler)

	getProductsHandler := kithttp.NewServer(
		makeGetProductsEndPoint(s),
		getProductsRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/paginated", getProductsHandler)

	addProductHandler := kithttp.NewServer(
		makeAddProductEndpoint(s),
		addProductRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/", addProductHandler)

	updateProductHandler := kithttp.NewServer(
		makeUpdateProductEndpoint(s),
		updateRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPut, "/", updateProductHandler)

	deleteProductHandler := kithttp.NewServer(
		makeDeleteProductEndpoint(s),
		deleteRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodDelete, "/{id}", deleteProductHandler)

	return r

}

func getProductByIdRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	//productId, _ := strconv.Atoi(chi.URLParam(r, "id"))
	productId := chi.URLParam(r, "id")
	return getProductByIDRequest{
		ProductID: productId,
	}, nil
}

func getProductsRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := getProductsRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}
	return request, nil
}

func addProductRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := getAddProductRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}
	return request, nil
}

func updateRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := updateProductRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}
	return request, nil
}

func deleteRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	productId := chi.URLParam(r, "id")
	return deleteProductRequest{
		ProductID: productId,
	}, nil
}
