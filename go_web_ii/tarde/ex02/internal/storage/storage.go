package storage

import (
	"errors"

	"github.com/gabrielobarbieri/api-meli/internal/model"
)

var (
	ErrStorageProductInternal = errors.New("storage: internal error")

	ErrStorageProductNotFound = errors.New("storage: product not found")

	ErrStorageProductInvalid = errors.New("storage: product invalid")
)

type StorageProduct interface {
	Get() (p []*model.Product, err error)
	GetByID(id int) (p *model.Product, err error)
	Search(price float64) (p []*model.Product, err error)
	Create(p *model.Product) (err error)
	UpdateOrCreate(p *model.Product) (err error)
	Update(p *model.Product) (err error)
	Delete(id int) (err error)
}
