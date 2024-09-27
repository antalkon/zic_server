package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Запуск HTTP-сервера в тестовом режиме
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	handler.ServeHTTP(w, req)

	if got := w.Body.String(); got != "Hello, World!" {
		t.Errorf("handler returned unexpected body: got %v want %v", got, "Hello, World!")
	}

	if got := w.Code; got != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", got, http.StatusOK)
	}
}
