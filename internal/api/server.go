package api

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Address string
	Port    uint16
}

func (s *Server) Run() error {
	http.HandleFunc("/", toHandler(testFunc))
	err := http.ListenAndServe(":8000", nil)
	return err
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func toHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			log.Fatal("Error:", err)
		}
	}
}

func testFunc(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("TESTED!")
	return nil
}
