package main

import (
	"fmt"
	transaction "protovalidate-demo/gen/ecommerce/transaction/v1"
	"testing"
	"time"

	"github.com/bufbuild/protovalidate-go"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Implement the ProtoMessage interface for Transaction
func (t *Transaction) Reset()         { *t = Transaction{} }
func (t *Transaction) String() string { return fmt.Sprintf("%v", *t) }
func (*Transaction) ProtoMessage()    {}

func TestValidate_ReturnsNoValidationErrors(t *testing.T) {
	// Create a Validator instance
	v, err := protovalidate.New()
	if err != nil {
		t.Fatalf("failed to create validator: %v", err)
	}

	// Create a valid Transaction message
	validTransaction := &transaction.Transaction{
		Id:           1000,
		Price:        "$100.00",
		PurchaseDate: timestamppb.New(time.Now().Add(-24 * time.Hour)), // 24 hours ago
		DeliveryDate: timestamppb.New(time.Now()),
		Email:        "example@example.com",
	}

	// Test valid transaction
	err = validate(v, validTransaction)
	assert.NoError(t, err, "expected no error for valid transaction")
}

func TestValidate_ReturnsValidationErrors(t *testing.T) {
	// Create a Validator instance
	v, err := protovalidate.New()
	if err != nil {
		t.Fatalf("failed to create validator: %v", err)
	}

	// Create an invalid Transaction message (e.g., missing required fields)
	invalidTransaction := &transaction.Transaction{
		Id:           999, // Invalid ID
		Price:        "",
		PurchaseDate: nil,
		DeliveryDate: nil,
		Email:        "",
	}

	// Test invalid transaction
	err = validate(v, invalidTransaction)
	assert.Error(t, err, "expected error for invalid transaction")
	assert.Contains(t, err.Error(), "validation failed", "expected validation failed error")
}
