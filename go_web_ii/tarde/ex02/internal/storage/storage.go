package storage

import (
	"errors"
	"time"
)

var (
	ErrStorageProductInternal = errors.New("storage: internal error")

	ErrStorageProductNotFound = errors.New("storage: product not found")

	ErrStorageProductInvalid = errors.New("storage: product invalid")
)

type Product struct {
	ID          int
	Name        string
	Quantity    int
	CodeValue   string
	IsPublished bool
	Expiration  time.Time
	Price       float64
}

type StorageProduct interface {
	Get() (p []*Product, err error)
	GetByID(id int) (p *Product, err error)
	Search(price float64) (p []*Product, err error)
	Create(p *Product) (err error)
}
