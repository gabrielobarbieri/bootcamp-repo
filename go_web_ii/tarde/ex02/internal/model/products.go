package model

import "time"

type Product struct {
	ID          int
	Name        string
	Quantity    int
	CodeValue   string
	IsPublished bool
	Expiration  time.Time
	Price       float64
}
