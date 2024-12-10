package storage

import (
	"fmt"
	"time"

	"github.com/gabrielobarbieri/api-meli/internal/model"
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

func (s *StorageProductMap) Create(p *model.Product) (err error) {
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

func (s *StorageProductMap) Get() ([]*model.Product, error) {
	products := make([]*model.Product, 0, len(s.DB))

	for k, v := range s.DB {
		products = append(products, &model.Product{ID: k, Name: v.Name, Quantity: v.Quantity, CodeValue: v.CodeValue, IsPublished: v.IsPublished, Expiration: v.Expiration, Price: v.Price})
	}

	return products, nil
}

func (s *StorageProductMap) GetByID(id int) (*model.Product, error) {
	product, ok := s.DB[id]

	if !ok {
		return nil, fmt.Errorf("%w: %d", ErrStorageProductNotFound, id)
	}

	p := &model.Product{ID: id, Name: product.Name, Quantity: product.Quantity, CodeValue: product.CodeValue, IsPublished: product.IsPublished, Expiration: product.Expiration, Price: product.Price}

	return p, nil
}

func (s *StorageProductMap) Search(priceGt float64) ([]*model.Product, error) {
	var products []*model.Product

	for k, v := range s.DB {
		if v.Price > priceGt {
			products = append(products, &model.Product{ID: k, Name: v.Name, Quantity: v.Quantity, CodeValue: v.CodeValue, IsPublished: v.IsPublished, Expiration: v.Expiration, Price: v.Price})
		}
	}

	return products, nil
}

func (s *StorageProductMap) UpdateOrCreate(p *model.Product) error {
	attr := &ProductAttributesMap{
		Name:        p.Name,
		Quantity:    p.Quantity,
		CodeValue:   p.CodeValue,
		IsPublished: p.IsPublished,
		Expiration:  p.Expiration,
		Price:       p.Price,
	}

	_, ok := s.DB[p.ID]

	switch ok {
	case true:
		s.DB[p.ID] = attr

	case false:
		s.LastId = len(s.DB) + 1
		s.DB[s.LastId] = attr
	}

	return nil
}

func (s *StorageProductMap) Update(p *model.Product) error {
	attr := &ProductAttributesMap{
		Name:        p.Name,
		Quantity:    p.Quantity,
		CodeValue:   p.CodeValue,
		IsPublished: p.IsPublished,
		Expiration:  p.Expiration,
		Price:       p.Price,
	}

	_, ok := s.DB[p.ID]

	switch ok {
	case true:
		s.DB[p.ID] = attr
	case false:
		return fmt.Errorf("%w: product not found", ErrStorageProductNotFound)
	}

	return nil
}

func (s *StorageProductMap) Delete(id int) error {
	_, ok := s.DB[id]

	switch ok {
	case true:
		delete(s.DB, id)
	case false:
		return fmt.Errorf("%w: product not found", ErrStorageProductNotFound)
	}

	return nil
}
