package storage

import (
	"fmt"
	"regexp"
)

type ConfigStorageProductValidate struct {
	St              StorageProduct
	RegexExpiration string
}

type StorageProductValidate struct {
	St              StorageProduct
	RegexExpiration *regexp.Regexp
}

func NewStorageProductValidate(cfg ConfigStorageProductValidate) *StorageProductValidate {
	if cfg.RegexExpiration == "" {
		cfg.RegexExpiration = `^\d{2}/\d{2}/\d{4}$`
	}

	return &StorageProductValidate{
		St:              cfg.St,
		RegexExpiration: regexp.MustCompile(cfg.RegexExpiration),
	}
}

func (s *StorageProductValidate) Create(p *Product) error {

	if p.Name == "" {
		return fmt.Errorf("%w: name is empty", ErrStorageProductInvalid)
	}

	if p.Quantity <= 0 {
		return fmt.Errorf("%w: quantity can't be zero or negative", ErrStorageProductInvalid)
	}

	if p.CodeValue == "" {
		return fmt.Errorf("%w: code value can't be empty", ErrStorageProductInvalid)
	}

	expirationStr := p.Expiration.Format("02/01/2006")

	if !s.RegexExpiration.MatchString(expirationStr) {
		return fmt.Errorf("%w: expiration format is invalid", ErrStorageProductInvalid)
	}

	if p.Price <= 0 {
		return fmt.Errorf("%w: price can't be zero or negative", ErrStorageProductInvalid)
	}

	err := s.St.Create(p)
	if err != nil {
		return err
	}
	return nil
}

func (s *StorageProductValidate) Get() ([]*Product, error) {
	p, err := s.St.Get()

	return p, err
}

func (s *StorageProductValidate) Search(priceGt float64) ([]*Product, error) {
	p, err := s.St.Search(priceGt)

	return p, err
}

func (s *StorageProductValidate) GetByID(id int) (*Product, error) {
	p, err := s.St.GetByID(id)

	return p, err
}
