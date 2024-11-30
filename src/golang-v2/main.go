package main

import (
	"fmt"
	"net/http"
	transaction "protovalidate-demo/gen/ecommerence/transaction/v1"
	hellov1 "protovalidate-demo/gen/example/hello/v1"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/protobuf/reflect/protoreflect"
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
		id:            1000,
		purchase_date: "2021-09-01",
		delivery_date: "2021-09-10",
		price:         "$5.99",
		email:         "bob@doe.mail",
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
