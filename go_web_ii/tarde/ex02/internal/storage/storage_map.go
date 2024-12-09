package storage

import (
	"fmt"
	"time"
)

type ProductAttributesMap struct {
	Name        string
	Quantity    int
	CodeValue   string
	IsPublished bool
	Expiration  time.Time
	Price       float64
}

type StorageProductMap struct {
	DB     map[int]*ProductAttributesMap
	LastId int
}

func NewStorageProductMap(DB map[int]*ProductAttributesMap, LastId int) *StorageProductMap {
	return &StorageProductMap{
		DB:     DB,
		LastId: LastId,
	}
}

func (s *StorageProductMap) Create(p *Product) (err error) {
	s.LastId = len(s.DB) + 1

	attr := ProductAttributesMap{
		Name:        p.Name,
		Quantity:    p.Quantity,
		CodeValue:   p.CodeValue,
		IsPublished: p.IsPublished,
		Expiration:  p.Expiration,
		Price:       p.Price,
	}

	s.DB[s.LastId] = &attr

	return nil
}

func (s *StorageProductMap) Get() ([]*Product, error) {
	products := make([]*Product, 0, len(s.DB))

	for k, v := range s.DB {
		products = append(products, &Product{k, v.Name, v.Quantity, v.CodeValue, v.IsPublished, v.Expiration, v.Price})
	}

	return products, nil
}

func (s *StorageProductMap) GetByID(id int) (*Product, error) {
	product, ok := s.DB[id]

	if !ok {
		return nil, fmt.Errorf("%w: %d", ErrStorageProductNotFound, id)
	}

	p := &Product{id, product.Name, product.Quantity, product.CodeValue, product.IsPublished, product.Expiration, product.Price}

	return p, nil
}

func (s *StorageProductMap) Search(priceGt float64) ([]*Product, error) {
	var products []*Product

	for k, v := range s.DB {
		if v.Price > priceGt {
			products = append(products, &Product{k, v.Name, v.Quantity, v.CodeValue, v.IsPublished, v.Expiration, v.Price})
		}
	}

	return products, nil
}
