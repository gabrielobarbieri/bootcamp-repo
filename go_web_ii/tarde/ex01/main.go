package main

import (
  "net/http"
  "fmt"
  "encoding/json"
  "os"
  "strconv"
  "errors"
  "regexp"
  "time"

  "github.com/go-chi/chi/v5"
)

type Product struct {
  ID int `json:"id"`
  Name string `json:"name"`
  Quantity int `json:"quantity"`
  CodeValue string `json:"codeValue"`
  IsPublished bool `json:"isPublished, omitempty"`
  Expiration string `json:"expiration"`
  Price float64 `json:"price"`
}

func loadProducts(filename string) ([]Product, error) {
  file, err := os.Open(filename)
  if err != nil {
    fmt.Println(err)
    return nil, errors.New("Could not find the file given!")
  }
  defer file.Close()

  products := []Product{}

  decoder := json.NewDecoder(file)
  err = decoder.Decode(&products)
  if err != nil {
    fmt.Println(err)
    return nil, errors.New("Could not get the content of the file!")
  }

  return products, nil
}

func handlePing(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodGet {
    w.WriteHeader(http.StatusMethodNotAllowed)
    fmt.Fprintln(w, "Method now allowed")
    return
  }

  w.WriteHeader(http.StatusOK)
  fmt.Fprintln(w, "pong")
}

func handleProductsList(w http.ResponseWriter, r *http.Request, products []Product) {
  if r.Method != http.MethodGet {
    w.WriteHeader(http.StatusMethodNotAllowed)
    fmt.Fprintln(w, "Method now allowed")
  }

  w.Header().Set("Content-Type", "application/json")
  encoder := json.NewEncoder(w)
  encoder.Encode(products)
}

func handleProducts(w http.ResponseWriter, r *http.Request, products []Product) {
  productIDStr := chi.URLParam(r, "id")
  productID, err := strconv.Atoi(productIDStr)
  if err != nil {
    http.Error(w, "Invalid ID", http.StatusBadRequest)
    return
  }

  for _, product := range products {
    if productID == product.ID {
      w.Header().Set("Content-Type", "application/json")
      encoder := json.NewEncoder(w)
      encoder.Encode(product)
      return
    }
  }

  http.Error(w, "No product found with this ID!", http.StatusNotFound)
}


func handleProductsSearch(w http.ResponseWriter, r *http.Request, products []Product) {
   //Take the param -> convert it to float64, -> search for products
  productPriceGTStr := r.URL.Query().Get("priceGt")
  productPriceGT, err := strconv.ParseFloat(productPriceGTStr, 64)
  if err != nil {
    http.Error(w, "Invalid price!", http.StatusBadRequest)
    return
  }

  productsResponse := []Product{}

  for _, product := range products {
    if product.Price > productPriceGT {
      productsResponse = append(productsResponse, product)
    }
  }

  w.Header().Set("Content-Type", "application/json")
  encoder := json.NewEncoder(w)
  err = encoder.Encode(productsResponse)
  if err != nil {
    http.Error(w, "Failed to encode products", http.StatusInternalServerError)
    return
  }
}

func handleProductCreation(w http.ResponseWriter, r *http.Request, p *[]Product) {
  if r.Method != http.MethodPost {
    http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
    return
  }

  // get the product model -> get info from req body -> populate new product -> put into the products list (value)
  var newProduct Product

  decoder := json.NewDecoder(r.Body)
  err := decoder.Decode(&newProduct)

  if err != nil {
    http.Error(w, "Could not decode the request body", http.StatusBadRequest)
    return
  }

  if newProduct.Name == "" || newProduct.CodeValue == "" || newProduct.Expiration == "" || newProduct.Quantity <= 0 || newProduct.Price <= 0 {
        http.Error(w, "All fields must be filled!", http.StatusBadRequest)
        return
    }

  for _, product := range *p {
    if product.CodeValue == newProduct.CodeValue {
      http.Error(w, "CodeValue must be unique", http.StatusBadRequest)
      return
    }
  }

  matched, _ := regexp.MatchString(`^\d{2}/\d{2}/\d{4}$`, newProduct.Expiration)

  if !matched {
    http.Error(w, "Expiration format must be XX/XX/XXXX", http.StatusBadRequest)
    return
  }

  _, err = time.Parse("02/01/2006", newProduct.Expiration)
  if err != nil {
    http.Error(w, "Expiration date is not valid", http.StatusBadRequest)
    return
  }

  newID := (*p)[len(*p) - 1].ID + 1
  newProduct.ID = newID

  *p = append(*p, newProduct)

  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(newProduct)
}

func main() {
  products, _ := loadProducts("./products.json")

  r := chi.NewRouter()

  r.Get("/ping", handlePing)

  r.Get("/products", func (w http.ResponseWriter, r *http.Request) {
    handleProductsList(w, r, products)
  })

  r.Get("/products/{id}", func (w http.ResponseWriter, r *http.Request) {
    handleProducts(w, r, products)
  })

  r.Get("/products/search", func (w http.ResponseWriter, r *http.Request) {
    handleProductsSearch(w, r, products)
  })

  r.Post("/products", func(w http.ResponseWriter, r *http.Request) {
    handleProductCreation(w, r, &products)
  })

  fmt.Println("Listening on port :8080")

  err := http.ListenAndServe(":8080", r)
  if err != nil {
    fmt.Println("Error starting server:", err)
    return
  }

}
