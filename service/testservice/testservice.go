package testservice

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

type Servicer interface {
	AllGroceries(w http.ResponseWriter, r Request)
}

func (s *Service) AllGroceries(w http.ResponseWriter, r Request) {

	fmt.Println(r)

	var groceries = []Grocery{
		{Name: "Almod Milk", Quantity: 2},
		{Name: "Apple", Quantity: 6},
	}

	fmt.Println("Endpoint hit: returnAllGroceries")
	json.NewEncoder(w).Encode(groceries)
}
