package handler

import (
	"bidfood/internal/model"
	"bidfood/internal/service"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestController_UpdateProductByID(t *testing.T) {
	t.Run("Update an existing product", func(t *testing.T) {

		productService := service.NewService()
		c := NewController(productService)
		c.ProductService.InsertProduct(model.Product{
			Name: "Car",
			Desc: "Toyota",
		})

		updateRequestProduct := model.Product{
			Id:   1, // Keep for comparing result
			Name: "Car",
			Desc: "Honda",
		}

		// Marshal the payload into JSON
		jsonPayload, err := json.Marshal(updateRequestProduct)
		if err != nil {
			t.Fatalf("Failed to marshal JSON payload: %v", err)
		}

		router := httprouter.New()
		router.PUT("/products/:id", c.UpdateProductByID)

		request, _ := http.NewRequest(http.MethodPut, "/products/1", bytes.NewBuffer(jsonPayload))
		t.Log(request)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		t.Log(response.Result().Status)
		var updatedResponseProduct model.Product
		json.Unmarshal(response.Body.Bytes(), &updatedResponseProduct)
		got := updatedResponseProduct
		want := updateRequestProduct

		t.Log("got ", got, " want", want)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
