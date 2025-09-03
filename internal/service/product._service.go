package service

import (
	"bidfood/internal/constant"
	"bidfood/internal/model"
	"errors"
	"log"
	"strings"
	"sync"
)

type Services interface {
	InsertProduct(product model.Product) (model.Product, error)
	GetProductByID(primaryID int) (*model.Product, error)
	UpdateProduct(product model.Product) (*model.Product, error)
	DeleteProduct(primaryID int) (int, error)
	GetAllProduct() ([]model.Product, error)
	GetPaginationProduct(page int) ([]model.Product, error)
	GetAllProductByFilter(filter string) ([]model.Product, error)
}

type ProductService struct {
	mu  sync.Mutex
	m   map[int]model.Product
	ids []int
	id  int
}

func NewService() *ProductService {
	return &ProductService{
		m:   make(map[int]model.Product),
		ids: make([]int, 0),
		id:  1,
	}
}

func (s *ProductService) InsertProduct(product model.Product) (model.Product, error) {
	s.mu.Lock()
	defer s.mu.Unlock() // Kayaknya ga perlu karna kan ini insert, bukan

	for _, existingProduct := range s.m {
		if existingProduct.Name == product.Name {
			return product, errors.New(constant.ProductAlreadyExist)
		}
	}

	s.ids = append(s.ids, s.id)
	product.Id = s.id
	s.m[s.id] = product
	s.id++

	return product, nil
}

func (s *ProductService) GetProductByID(primaryID int) (*model.Product, error) {
	val, ok := s.m[primaryID]
	if ok {
		return &val, nil
	}
	return nil, errors.New(constant.ProductNotFoundErr)
}

func (s *ProductService) UpdateProduct(product model.Product) (*model.Product, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.m[product.Id]; exists {
		s.m[product.Id] = product
		return &product, nil
	}
	// Belum ada penjagaan kl ID nya ga ada
	return &product, errors.New(constant.ProductNotFoundErr)
}

func (s *ProductService) DeleteProduct(primaryID int) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.m[primaryID]; exists {
		delete(s.m, primaryID)
		s.ids = append(s.ids[:primaryID-1], s.ids[primaryID:]...)
		log.Printf("Product with id %d was deleted.", primaryID)
		return primaryID, nil
	}

	log.Printf("Product with id %d not exist.", primaryID)
	return primaryID, errors.New(constant.ProductNotFoundErr)
}

func (s *ProductService) GetAllProduct() ([]model.Product, error) {

	var result []model.Product
	result = make([]model.Product, 0)

	log.Println(">>", s)

	if len(s.ids) == 0 {
		return nil, errors.New(constant.EmptyResultErr)
	}
	for _, id := range s.ids {

		log.Printf(">> GetAllProduct add : %d - %v", id, s.m[id])
		result = append(result, s.m[id])
	}
	// Belum ada penjagaan kl ID nya ga ada
	return result, nil
}

func (s *ProductService) GetAllProductByFilter(filterName string) ([]model.Product, error) {
	var result []model.Product
	result = make([]model.Product, 0)

	log.Println(">>", s)

	if len(s.ids) == 0 {
		return nil, errors.New(constant.EmptyResultErr)
	}
	for _, id := range s.ids {

		log.Printf(">> GetAllProduct add : %d - %v", id, s.m[id])

		if strings.Contains(strings.ToLower(s.m[id].Name), (strings.ToLower(filterName))) {
			result = append(result, s.m[id])
		}
	}

	if len(result) == 0 {
		return nil, errors.New(constant.EmptyResultErr)
	}
	// Belum ada penjagaan kl ID nya ga ada
	return result, nil
}

func (s *ProductService) GetPaginationProduct(page int) ([]model.Product, error) {

	var itemsInPage = 5
	startIndex := page * itemsInPage
	if startIndex >= len(s.ids) {
		return nil, errors.New(constant.ExceedsRequest)
	}

	var result []model.Product
	startID := s.ids[startIndex]

	for i, size := 0, 0; startID+i < s.id && size < itemsInPage; i++ {
		if product, ok := s.m[startID+i]; ok {
			go func() {
				log.Println(product)
			}()
			result = append(result, product)
			size++
		}
	}

	// Belum ada penjagaan kl pagenya ketinggian
	return result, nil
}
