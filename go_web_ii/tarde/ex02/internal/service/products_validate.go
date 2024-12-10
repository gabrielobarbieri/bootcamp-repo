package service

import (
	"fmt"
	"regexp"

	"github.com/gabrielobarbieri/api-meli/internal/model"
	"github.com/gabrielobarbieri/api-meli/internal/storage"
)

type ConfigStorageProductValidate struct {
	St              storage.StorageProduct
	RegexExpiration string
}

type StorageProductValidate struct {
	St              storage.StorageProduct
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

func (s *StorageProductValidate) Create(p *model.Product) error {

	if p.Name == "" {
		return fmt.Errorf("%w: name is empty", ErrValidatorProductFieldRequired)
	}

	if p.Quantity <= 0 {
		return fmt.Errorf("%w: quantity can't be zero or negative", ErrValidatorProductFieldRequired)
	}

	if p.CodeValue == "" {
		return fmt.Errorf("%w: code value can't be empty", ErrValidatorProductFieldRequired)
	}

	expirationStr := p.Expiration.Format("02/01/2006")

	if !s.RegexExpiration.MatchString(expirationStr) {
		return fmt.Errorf("%w: expiration format is invalid", storage.ErrStorageProductInvalid)
	}

	if p.Price <= 0 {
		return fmt.Errorf("%w: price can't be zero or negative", storage.ErrStorageProductInvalid)
	}

	err := s.St.Create(p)
	if err != nil {
		return err
	}
	return nil
}

func (s *StorageProductValidate) Get() ([]*model.Product, error) {
	p, err := s.St.Get()

	return p, err
}

func (s *StorageProductValidate) Search(priceGt float64) ([]*model.Product, error) {
	p, err := s.St.Search(priceGt)

	return p, err
}

func (s *StorageProductValidate) GetByID(id int) (*model.Product, error) {
	p, err := s.St.GetByID(id)

	return p, err
}

func (s *StorageProductValidate) UpdateOrCreate(p *model.Product) error {
	if p.Name == "" {
		return fmt.Errorf("%w: name is empty", ErrValidatorProductFieldRequired)
	}

	if p.Quantity <= 0 {
		return fmt.Errorf("%w: quantity can't be zero or negative", ErrValidatorProductFieldInvalid)
	}

	if p.CodeValue == "" {
		return fmt.Errorf("%w: code value can't be empty", ErrValidatorProductFieldRequired)
	}

	expirationStr := p.Expiration.Format("02/01/2006")

	if !s.RegexExpiration.MatchString(expirationStr) {
		return fmt.Errorf("%w: expiration format is invalid", ErrValidatorProductFieldInvalid)
	}

	if p.Price <= 0 {
		return fmt.Errorf("%w: price can't be zero or negative", ErrValidatorProductFieldInvalid)
	}

	err := s.St.UpdateOrCreate(p)
	if err != nil {
		return err
	}

	return nil
}

func (s *StorageProductValidate) Update(p *model.Product) error {
	err := s.St.Update(p)
	return err
}

func (s *StorageProductValidate) Delete(id int) error {
	err := s.St.Delete(id)

	return err
}
