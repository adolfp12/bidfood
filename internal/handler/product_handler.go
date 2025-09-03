package handler

import (
	"bidfood/internal/model"
	"bidfood/internal/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Controller struct {
	ProductService service.Services
}

// func New() *Controller {
// 	return &Controller{
// 		ProductService: New(),
// 	}
// }

func NewController(service service.Services) *Controller {
	return &Controller{
		ProductService: service,
	}
}

func (c *Controller) GetAllProduct(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var resp []model.Product
	var err error
	if idStr := strings.TrimSpace(r.URL.Query().Get("page")); idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid page number", http.StatusBadRequest)
			return
		}
		id = id - 1
		if id < 0 {
			http.Error(w, "Invalid page number", http.StatusBadRequest)
			return
		}
		resp, err = c.ProductService.GetPaginationProduct(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else if filter := strings.TrimSpace(r.URL.Query().Get("filter")); filter != "" {
		resp, err = c.ProductService.GetAllProductByFilter(filter)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		resp, err = c.ProductService.GetAllProduct()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (c *Controller) AddNewProduct(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var product model.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	resp, err := c.ProductService.InsertProduct(product)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set header response ke application/json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// Encode struct ke JSON dan kirim ke client
	json.NewEncoder(w).Encode(resp)

}

func (c *Controller) GetProductByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idStr := params.ByName("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid converting id", http.StatusBadRequest)
		return
	}

	resp, err := c.ProductService.GetProductByID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set header response ke application/json
	w.Header().Set("Content-Type", "application/json")
	// Encode struct ke JSON dan kirim ke client
	json.NewEncoder(w).Encode(resp)

}

/*
Update product by ID. Example http://localhost:8080/1
*/
func (c *Controller) UpdateProductByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	idStr := params.ByName("id")
	id, err := strconv.Atoi(idStr)
	log.Printf("idStr: %s; id:%d", idStr, id)
	if err != nil {
		http.Error(w, "Invalid converting id", http.StatusBadRequest)
		return
	}

	var product model.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	log.Printf("product: %v", product)

	product.Id = id

	log.Printf("product with id: %v", product)
	resp, err := c.ProductService.UpdateProduct(product)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set header response ke application/json
	w.Header().Set("Content-Type", "application/json")
	// Encode struct ke JSON dan kirim ke client
	json.NewEncoder(w).Encode(resp)
}

// handle delete product by ID
func (c *Controller) DeleteProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idStr := params.ByName("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	resp, err := c.ProductService.DeleteProduct(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set header response ke application/json
	w.Header().Set("Content-Type", "application/json")
	// Encode struct ke JSON dan kirim ke client
	json.NewEncoder(w).Encode(resp)
}

func (c *Controller) TestAddProduct(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	productsTest := []model.Product{
		{Name: "Apple", Desc: "Red"},
		{Name: "Banana", Desc: "Yellow"},
		{Name: "Orange", Desc: "Orange"},
		{Name: "Strawberry", Desc: "Red"},
		{Name: "Mango", Desc: "Green"},
		{Name: "Grapes", Desc: "Red"},
		{Name: "Pineapple", Desc: "Yellow"},
		{Name: "Blueberry", Desc: "Blue to Purple"},
		{Name: "Watermelon", Desc: "Green"},
		{Name: "Papaya", Desc: "Orange"},
		{Name: "Kiwi", Desc: "Green"},
		{Name: "Peach", Desc: "Orange"},
		{Name: "Cherry", Desc: "Purple"},
		{Name: "Pomegranate", Desc: "Red"},
		{Name: "Dragon Fruit", Desc: "Pink"},
	}
	for _, fruit := range productsTest {
		product := model.Product{Name: fruit.Name, Desc: fruit.Desc}
		c.ProductService.InsertProduct(product)
	}

	w.WriteHeader(http.StatusOK) // optional, default 200 OK

	fmt.Fprint(w, "OK - Insert!!")
}

func (c *Controller) Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	time.Sleep(10 * time.Second)
	w.Write([]byte("Home\n"))
}
