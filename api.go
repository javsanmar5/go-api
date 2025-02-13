package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application-json")
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			// Here is where we handle the error
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenADDR string
}

func NewAPIServer(listenADDR string) *APIServer {
	return &APIServer{
		listenADDR: listenADDR,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/example",  makeHTTPHandleFunc(s.handleExample))
	log.Println("Server is running on port: ", s.listenADDR)
	http.ListenAndServe(s.listenADDR, router)
}

func (s *APIServer) handleExample(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetExample(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateExample(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteExample(w, r)
	}

	return fmt.Errorf("HTTP Method not allowed: %s", r.Method) 
}

func (s *APIServer) handleGetExample(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleCreateExample(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteExample(w http.ResponseWriter, r *http.Request) error {
	return nil
}

