package loader

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gabrielobarbieri/api-meli/internal/storage"
)

type LoaderJSON struct {
	filepath string
}

func NewLoaderJSON(filepath string) *LoaderJSON {
	return &LoaderJSON{filepath}
}

type ProductAttributesJSON struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func (l *LoaderJSON) Load() (p *ProductsDB, err error) {
	f, err := os.Open(l.filepath)
	if err != nil {
		err = fmt.Errorf("%w: %s", ErrLoaderProductInternal, err)
		return
	}

	defer f.Close()
	// Create DB and read the file

	DB := make(map[int]*storage.ProductAttributesMap)

	products := []ProductAttributesJSON{}

	err = json.NewDecoder(f).Decode(&products)
	if err != nil {
		err = fmt.Errorf("%w: %s", ErrLoaderProductInternal, err)
		return
	}

	for _, product := range products {
		exp, err := time.Parse("02/01/2006", product.Expiration)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrLoaderProductInternal, err)
		}

		DB[product.ID] = &storage.ProductAttributesMap{
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  exp,
			Price:       product.Price,
		}
	}

	lastId := len(DB) + 1

	p = &ProductsDB{DB: DB, LastId: lastId}

	return
}
