package controller

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gabrielobarbieri/api-meli/internal/model"
	"github.com/gabrielobarbieri/api-meli/internal/service"
	"github.com/gabrielobarbieri/api-meli/internal/storage"
	"github.com/gabrielobarbieri/api-meli/pkg/web/request"
	"github.com/gabrielobarbieri/api-meli/pkg/web/response"

	"github.com/go-chi/chi"
)

type HandlerProducts struct {
	Sv service.Service
}

func NewHandlerProducts(sv service.Service) *HandlerProducts {
	return &HandlerProducts{Sv: sv}
}

type ProductJSON struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type ResponseBodyProductGet struct {
	Message string         `json:"message"`
	Data    []*ProductJSON `json:"data"`
}

func (h *HandlerProducts) Get(w http.ResponseWriter, r *http.Request) {
	// Create Response body -> call pkg method
	p, err := h.Sv.Get()
	if err != nil {
		code := http.StatusInternalServerError
		body := ResponseBodyProductGet{Message: "internal error", Data: nil}

		response.JSON(w, code, body)
		return
	}

	products := make([]*ProductJSON, 0, len(p))

	for _, v := range p {
		product := &ProductJSON{v.ID, v.Name, v.Quantity, v.CodeValue, v.IsPublished, v.Expiration.Format("02/01/2006"), v.Price}
		products = append(products, product)
	}

	code := http.StatusOK
	body := ResponseBodyProductGet{Message: "Products", Data: products}

	response.JSON(w, code, body)
}

type ResponseBodyGetById struct {
	Message string       `json:"message"`
	Data    *ProductJSON `json:"data"`
}

func (h *HandlerProducts) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		body := ResponseBodyGetById{Message: "invalid id", Data: nil}
		code := http.StatusBadRequest

		response.JSON(w, code, body)
		return
	}

	p, err := h.Sv.GetByID(id)
	if err != nil {
		var code int
		var body ResponseBodyGetById

		switch {
		case errors.Is(err, storage.ErrStorageProductNotFound):
			code = http.StatusNotFound
			body = ResponseBodyGetById{
				Message: "product not found",
				Data:    nil,
			}
		default:
			code = http.StatusInternalServerError
			body = ResponseBodyGetById{
				Message: "internal server error",
				Data:    nil,
			}
		}

		response.JSON(w, code, body)
		return
	}

	code := http.StatusOK
	body := ResponseBodyGetById{
		Message: "Product",
		Data:    &ProductJSON{p.ID, p.Name, p.Quantity, p.CodeValue, p.IsPublished, p.Expiration.Format("02/01/2006"), p.Price},
	}

	response.JSON(w, code, body)
}

type ResponseBodySearch struct {
	Message string         `json:"message"`
	Data    []*ProductJSON `json:"data"`
}

func (h *HandlerProducts) Search(w http.ResponseWriter, r *http.Request) {
	priceGt, err := strconv.ParseFloat(r.URL.Query().Get("priceGt"), 64)
	if err != nil {
		code := http.StatusBadRequest
		body := ResponseBodySearch{
			Message: "invalid price",
			Data:    nil,
		}

		response.JSON(w, code, body)
		return
	}

	p, err := h.Sv.Search(priceGt)
	if err != nil {
		code := http.StatusInternalServerError
		body := ResponseBodySearch{
			Message: "internal server error",
			Data:    nil,
		}

		response.JSON(w, code, body)
		return
	}

	var products []*ProductJSON

	for _, product := range p {
		products = append(products, &ProductJSON{product.ID, product.Name, product.Quantity, product.CodeValue, product.IsPublished, product.Expiration.Format("02/01/2006"), product.Price})
	}

	code := http.StatusOK
	body := ResponseBodySearch{
		Message: "products",
		Data:    products,
	}

	response.JSON(w, code, body)
}

type ResponseBodyCreate struct {
	Message string       `json:"message"`
	Data    *ProductJSON `json:"data"`
}

type RequestBodyCreate struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func (h *HandlerProducts) Create(w http.ResponseWriter, r *http.Request) {
	p := &RequestBodyCreate{}

	err := request.JSON(r, p)
	if err != nil {
		code := http.StatusBadRequest
		body := ResponseBodyCreate{
			Message: "invalid request body",
			Data:    nil,
		}

		response.JSON(w, code, body)
		return
	}

	exp, err := time.Parse("02/01/2006", p.Expiration)
	if err != nil {
		code := http.StatusBadRequest
		body := ResponseBodyCreate{
			Message: "Invalid time format",
			Data:    nil,
		}

		response.JSON(w, code, body)
		return
	}

	product := &model.Product{
		Name:        p.Name,
		Quantity:    p.Quantity,
		CodeValue:   p.CodeValue,
		IsPublished: p.IsPublished,
		Expiration:  exp,
		Price:       p.Price,
	}

	err = h.Sv.Create(product)
	if err != nil {
		var code int
		var body ResponseBodyCreate

		switch {
		case errors.Is(err, storage.ErrStorageProductInvalid):
			code = http.StatusUnprocessableEntity
			body = ResponseBodyCreate{
				Message: "invalid product",
				Data:    nil,
			}
		default:
			code = http.StatusBadRequest
			body = ResponseBodyCreate{
				Message: "internal server error",
				Data:    nil,
			}
		}

		response.JSON(w, code, body)
		return
	}

	code := http.StatusCreated
	body := ResponseBodyCreate{
		Message: "product created",
		Data: &ProductJSON{
			product.ID,
			product.Name,
			product.Quantity,
			product.CodeValue,
			product.IsPublished,
			product.Expiration.Format("02/01/2006"),
			product.Price,
		},
	}

	response.JSON(w, code, body)
}

type ResponsePing struct {
	Message string
	Data    *string
}

func (h *HandlerProducts) Ping(w http.ResponseWriter, r *http.Request) {
	code := http.StatusOK
	body := ResponsePing{
		Message: "Pong",
		Data:    nil,
	}

	response.JSON(w, code, body)
}

type ResponseBodyUpdateOrCreate struct {
	Message string
	Data    *ProductJSON
}

type RequestBodyUpdateOrCreate struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func (h *HandlerProducts) UpdateOrCreate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		body := ResponseBodyGetById{Message: "invalid id", Data: nil}
		code := http.StatusBadRequest

		response.JSON(w, code, body)
		return
	}

	var reqBody RequestBodyUpdateOrCreate

	if err = request.JSON(r, &reqBody); err != nil {
		code := http.StatusBadRequest
		body := ResponseBodyUpdateOrCreate{
			Message: "invalid request body",
			Data:    nil,
		}

		response.JSON(w, code, body)
		return
	}

	exp, err := time.Parse("02/01/2006", reqBody.Expiration)
	if err != nil {
		code := http.StatusBadRequest
		body := ResponseBodyUpdateOrCreate{
			Message: "expiration format should be xx/xx/xxxx",
			Data:    nil,
		}

		response.JSON(w, code, body)
	}

	p := &model.Product{
		ID:          id,
		Name:        reqBody.Name,
		Quantity:    reqBody.Quantity,
		CodeValue:   reqBody.CodeValue,
		IsPublished: reqBody.IsPublished,
		Expiration:  exp,
		Price:       reqBody.Price,
	}

	err = h.Sv.UpdateOrCreate(p)
	if err != nil {
		var code int
		var body ResponseBodyCreate

		switch {
		case errors.Is(err, storage.ErrStorageProductInvalid):
			code = http.StatusUnprocessableEntity
			body = ResponseBodyCreate{
				Message: "invalid product",
				Data:    nil,
			}
		default:
			code = http.StatusBadRequest
			body = ResponseBodyCreate{
				Message: "internal server error",
				Data:    nil,
			}
		}

		response.JSON(w, code, body)
		return
	}

	code := http.StatusCreated
	body := ResponseBodyUpdateOrCreate{
		Message: "product created or updated",
		Data:    &ProductJSON{p.ID, p.Name, p.Quantity, p.CodeValue, p.IsPublished, p.Expiration.Format("2006-01-02"), p.Price},
	}

	response.JSON(w, code, body)
}

type RequestBodyUpdate struct {
	Name        string
	Quantity    int
	CodeValue   string
	IsPublished bool
	Expiration  string
	Price       float64
}

type ResponseBodyUpdate struct {
	Message string
	Data    *ProductJSON
}

func (h *HandlerProducts) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		body := ResponseBodyGetById{Message: "invalid id", Data: nil}
		code := http.StatusBadRequest

		response.JSON(w, code, body)
		return
	}

	var reqBody RequestBodyUpdate

	if err := request.JSON(r, &reqBody); err != nil {
		code := http.StatusBadRequest
		body := ResponseBodyUpdateOrCreate{
			Message: "invalid request body",
			Data:    nil,
		}

		response.JSON(w, code, body)
		return
	}

	exp, err := time.Parse("02/01/2006", reqBody.Expiration)
	if err != nil {
		code := http.StatusBadRequest
		body := ResponseBodyUpdateOrCreate{
			Message: "expiration format should be xx/xx/xxxx",
			Data:    nil,
		}

		response.JSON(w, code, body)
	}

	p := &model.Product{
		ID:          id,
		Name:        reqBody.Name,
		Quantity:    reqBody.Quantity,
		CodeValue:   reqBody.CodeValue,
		IsPublished: reqBody.IsPublished,
		Expiration:  exp,
		Price:       reqBody.Price,
	}

	err = h.Sv.Update(p)
	if err != nil {
		var code int
		var body ResponseBodyCreate

		switch {
		case errors.Is(err, storage.ErrStorageProductNotFound):
			code = http.StatusNotFound
			body = ResponseBodyCreate{
				Message: "product not found",
				Data:    nil,
			}
		default:
			code = http.StatusBadRequest
			body = ResponseBodyCreate{
				Message: "internal server error",
				Data:    nil,
			}
		}

		response.JSON(w, code, body)
		return
	}

	code := http.StatusOK
	body := ResponseBodyUpdate{
		Message: "product updated",
		Data:    nil,
	}

	response.JSON(w, code, body)
}

type ResponseBodyDelete struct {
	Message string
	Data    *ProductJSON
}

func (h *HandlerProducts) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		body := ResponseBodyGetById{Message: "invalid id", Data: nil}
		code := http.StatusBadRequest

		response.JSON(w, code, body)
		return
	}

	err = h.Sv.Delete(id)
	if err != nil {
		var code int
		var body ResponseBodyCreate

		switch {
		case errors.Is(err, storage.ErrStorageProductInvalid):
			code = http.StatusNotFound
			body = ResponseBodyCreate{
				Message: "product not found",
				Data:    nil,
			}
		default:
			code = http.StatusBadRequest
			body = ResponseBodyCreate{
				Message: "internal server error",
				Data:    nil,
			}
		}

		response.JSON(w, code, body)
		return
	}

	code := http.StatusNoContent
	body := ResponseBodyDelete{
		Message: "product deleted",
		Data:    nil,
	}

	response.JSON(w, code, body)
}
