package utils

import (
	"net/http"
	"testing"
)

func TestBadRequestError(t *testing.T) {
	status, err := BadRequestError()
	if status != http.StatusBadRequest {
		t.Errorf("Expected status %v, got %v", http.StatusBadRequest, status)
	}
	if err.Message != "Bad Request" {
		t.Errorf("Expected message 'Bad Request', got %v", err.Message)
	}
}

func TestNotFoundError(t *testing.T) {
	status, err := NotFoundError()
	if status != http.StatusNotFound {
		t.Errorf("Expected status %v, got %v", http.StatusNotFound, status)
	}
	if err.Message != "Not Found" {
		t.Errorf("Expected message 'Not Found', got %v", err.Message)
	}
}

func TestInternalServerError(t *testing.T) {
	msg := "Internal Server Error"
	status, err := InternalServerError(msg)
	if status != http.StatusInternalServerError {
		t.Errorf("Expected status %v, got %v", http.StatusInternalServerError, status)
	}
	if err.Message != msg {
		t.Errorf("Expected message %v, got %v", msg, err.Message)
	}
}

func TestConflictError(t *testing.T) {
	status, err := ConflictError()
	if status != http.StatusConflict {
		t.Errorf("Expected status %v, got %v", http.StatusConflict, status)
	}
	if err.Message != "Some fields can't be modified" {
		t.Errorf("Expected message 'Some fields can't be modified', got %v", err.Message)
	}
}

func TestMultipleLoginError(t *testing.T) {
	err := MultipleLoginError()
	if err.Message != "login should be unique" {
		t.Errorf("Expected message 'login should be unique', got %v", err.Message)
	}
}
