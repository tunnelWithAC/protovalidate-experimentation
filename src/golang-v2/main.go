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

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s", name)
}

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

func validateHandler(w http.ResponseWriter, r *http.Request) {
	v, err := protovalidate.New()
	if err != nil {
		panic(err)
	}

	name := r.URL.Query().Get("value")
	fmt.Println("Name: ", name)

	var errorList []error

	if err = validate(v, &transaction.Transaction{
		Id:           1000,
		Price:        "$100.00",
		PurchaseDate: timestamppb.New(time.Now().Add(-24 * time.Hour)), // 24 hours ago
		DeliveryDate: timestamppb.New(time.Now()),
		Email:        "bob@doe.mail",
	}); err != nil {
		errorList = append(errorList, err)
	}

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
		PurchaseDate: (*timestamppb.Timestamp)(&transactionRequest.PurchaseDate), // Convert to *timestamppb.Timestamp
		DeliveryDate: (*timestamppb.Timestamp)(&transactionRequest.DeliveryDate), // Convert to *timestamppb.Timestamp
		Email:        transactionRequest.Email,
	}); err != nil {
		errorList = append(errorList, err)
		http.Error(w, "Request failed validation", http.StatusBadRequest)
		return
	}

	// Respond with the created transaction
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactionRequest)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	http.HandleFunc("/validate", validateHandler)
	fmt.Println("Server is running on port 8080...")
	r.Post("/transactions", transactionHandler)

	http.ListenAndServe(":8080", r)
}
