package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCalculaeRetornaMensagem(t *testing.T) {
	obtido := calculaeRetornaMensagem("Teste")
	esperado := "<b>Teste</b>"

	if obtido == esperado {
		t.Errorf("calculaeRetornaMensagem('Teste wrong') \n obtido: %v \n esperado:  \n%v", obtido, esperado)
	}
}

func TestCalculaRetornaMensagemHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(calculaeRetornaMensagemHandler)

	handler.ServeHTTP(rr, req)

	// Check o status code é o esperado
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler retornou status code errado: obtido %v esperado %v",
			status, http.StatusOK)
	}

	// Check o response body é o esperado.
	esperado := "Code Education Rocks!"
	if rr.Body.String() != esperado {
		t.Errorf("handler conteúdo inesperado: obtido: %v esperado: %v",
			rr.Body.String(), esperado)
	}
}

func TestHandler404(t *testing.T) {

	req, err := http.NewRequest("GET", "/teste", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(calculaeRetornaMensagemHandler)

	handler.ServeHTTP(rr, req)

	// Check o status code é o esperado
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler retornou status code errado: obtido status: %v esperado status: %v",
			status, http.StatusNotFound)
	}

	esperado := "404 nao encontrado."
	if obtido := strings.TrimSpace(rr.Body.String()); obtido != esperado {
		t.Errorf("handler retornou mensagem de erro errada: obtida: %v esperada: %v",
			obtido, esperado)
	}
}

func TestHandlerPOST(t *testing.T) {

	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(calculaeRetornaMensagemHandler)

	handler.ServeHTTP(rr, req)

	// Check o status code é o esperado
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler retornou status code errado: obtido status: %v esperado status: %v",
			status, http.StatusNotFound)
	}

	esperado := "Metodo nao suportado."
	if obtido := strings.TrimSpace(rr.Body.String()); obtido != esperado {
		t.Errorf("handler retornou mensagem de erro errada: obtida: %v esperada: %v",
			obtido, esperado)
	}
}
