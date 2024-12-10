package service

import (
	"errors"

	"github.com/gabrielobarbieri/api-meli/internal/model"
)

var (
	ErrValidatorProductFieldRequired = errors.New("validator: product field required")

	ErrValidatorProductFieldInvalid = errors.New("validator: product field invalid")
)

type Service interface {
	Get() (p []*model.Product, err error)
	GetByID(id int) (p *model.Product, err error)
	Search(price float64) (p []*model.Product, err error)
	Create(p *model.Product) (err error)
	UpdateOrCreate(p *model.Product) (err error)
	Update(p *model.Product) (err error)
	Delete(id int) (err error)
}
