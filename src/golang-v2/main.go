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
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s", name)
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

func main() {
	http.HandleFunc("/validate", validateHandler)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
