package main

import (
	"fmt"
	"net/http"
	transaction "protovalidate-demo/gen/ecommerce/transaction/v1"
	hellov1 "protovalidate-demo/gen/example/hello/v1"
	"time"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"

	"encoding/json"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Timestamp timestamppb.Timestamp

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	parsedTime, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return err
	}
	*t = Timestamp(*timestamppb.New(parsedTime))
	return nil
}

type Transaction struct {
	Id           uint64    `json:"id"`
	Price        string    `json:"price"`
	PurchaseDate Timestamp `json:"purchase_date"`
	DeliveryDate Timestamp `json:"delivery_date"`
	Email        string    `json:"email"`
}

func validate(v *protovalidate.Validator, msg protoreflect.ProtoMessage) error {
	if err := v.Validate(msg); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	return nil
}

func nameHandler(w http.ResponseWriter, r *http.Request) {
	v, err := protovalidate.New()
	if err != nil {
		panic(err)
	}

	name := r.URL.Query().Get("value")
	fmt.Println("Name: ", name)

	var errorList []error

	if err = validate(v, &hellov1.Name{
		Name: name,
	}); err != nil {
		errorList = append(errorList, err)
	}

	fmt.Fprintf(w, "Errors:  %s", errorList)
}

func transactionHandler(w http.ResponseWriter, r *http.Request) {
	var transactionRequest Transaction
	err := json.NewDecoder(r.Body).Decode(&transactionRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var errorList []error
	v, err := protovalidate.New()
	if err != nil {
		panic(err)
	}

	if err = validate(v, &transaction.Transaction{
		Id:           transactionRequest.Id,
		Price:        transactionRequest.Price,
		PurchaseDate: (*timestamppb.Timestamp)(&transactionRequest.PurchaseDate),
		DeliveryDate: (*timestamppb.Timestamp)(&transactionRequest.DeliveryDate),
		Email:        transactionRequest.Email,
	}); err != nil {
		errorList = append(errorList, err)
		// Convert []error to []string
		var errorStrings []string
		for _, e := range errorList {
			errorStrings = append(errorStrings, e.Error())
		}

		// Marshal []string to JSON
		errorJSON, err := json.Marshal(errorStrings)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Error(w, string(errorJSON), http.StatusBadRequest)
		return
	}

	// Respond with the created transaction
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactionRequest)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	http.HandleFunc("/validate", nameHandler)
	r.Post("/transactions", transactionHandler)
	fmt.Println("Server is running on port 8080...")

	http.ListenAndServe(":8080", r)
}
