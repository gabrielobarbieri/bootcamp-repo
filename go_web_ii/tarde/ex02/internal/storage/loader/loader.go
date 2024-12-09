package loader

import (
	"errors"

	"github.com/gabrielobarbieri/api-meli/internal/storage"
)

var (
	ErrLoaderProductInternal = errors.New("Loader: internal error")
)

type ProductsDB struct {
	DB     map[int]*storage.ProductAttributesMap
	LastId int
}

type Loader interface {
	Load() (p *ProductsDB, err error)
}
