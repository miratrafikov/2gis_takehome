package controller

import (
	"applicationDesignTest/internal/model"
	createorder "applicationDesignTest/internal/usecases/create_order"
	"encoding/json"
	"net/http"
)

type CreateOrder struct {
	uc createorder.Usecase
}

func NewCreateOrder(uc createorder.Usecase) CreateOrder {
	return CreateOrder{uc: uc}
}

func (c CreateOrder) Handle(w http.ResponseWriter, r *http.Request) {
	var body model.Order
	json.NewDecoder(r.Body).Decode(&body)
	order, err := c.uc.Handle(body)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(order)
	}
}
