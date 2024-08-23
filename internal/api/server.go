package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	IP    string
	Port  uint16
	Store Store
}

func (s *Server) Run() error {
	router := chi.NewRouter()
	router.Post("/users", toHandler(s.createHandler))
	router.Get("/users", toHandler(s.getUsersHandler))
	router.Put("/users/{id}", toHandler(s.updateHandler))
	router.Get("/users/{id}", toHandler(s.getUserHandler))
	router.Delete("/users/{id}", toHandler(s.deleteHandler))
	router.Put("/heart/{id}", toHandler(s.heartHandler))

	fmt.Println("API SERVER RUNNING ON PORT", s.Port)
	apiServerAddr := fmt.Sprintf("%s:%d", s.IP, s.Port)
	err := http.ListenAndServe(apiServerAddr, router)
	return err
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type sensibleHandler func(w http.ResponseWriter, r *http.Request) (error, int)

func toHandler(f sensibleHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err, status := f(w, r); err != nil {
			WriteJSON(w, status, nil)
			log.Fatal("Error:", err)
		}
	}
}

func (s *Server) getUsersHandler(w http.ResponseWriter, r *http.Request) (error, int) {
	snackers, err := s.Store.FetchSnackers()
	if err != nil {
		return err, http.StatusInternalServerError
	}

	WriteJSON(w, http.StatusOK, snackers)
	return nil, http.StatusOK
}

func (s *Server) getUserHandler(w http.ResponseWriter, r *http.Request) (error, int) {
	id, err := parseURLID(r)
	if err != nil {
		return err, http.StatusBadRequest
	}
	snacker, err := s.Store.FindSnacker(id)
	WriteJSON(w, http.StatusOK, snacker)
	return nil, http.StatusOK
}

func (s *Server) createHandler(w http.ResponseWriter, r *http.Request) (error, int) {
	var snackerReg SnackerRegistration
	err := json.NewDecoder(r.Body).Decode(&snackerReg)
	if err != nil {
		return err, http.StatusBadRequest
	}

	id, err := s.Store.CreateSnacker(&snackerReg)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	foundSnacker, err := s.Store.FindSnacker(id)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	WriteJSON(w, http.StatusOK, foundSnacker)
	return nil, http.StatusOK
}

func (s *Server) updateHandler(w http.ResponseWriter, r *http.Request) (error, int) {
	var snacker Snacker
	err := json.NewDecoder(r.Body).Decode(&snacker)
	if err != nil {
		return err, http.StatusBadRequest
	}

	urlID, err := parseURLID(r)
	if err != nil {
		return err, http.StatusBadRequest
	} else if urlID != snacker.ID {
		return errors.New("Cannot modify user ID"), http.StatusBadRequest
	}

	err = s.Store.UpdateSnacker(&snacker)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	foundSnacker, err := s.Store.FindSnacker(snacker.ID)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	WriteJSON(w, http.StatusOK, foundSnacker)
	return nil, http.StatusOK
}

func (s *Server) deleteHandler(w http.ResponseWriter, r *http.Request) (error, int) {
	id, err := parseURLID(r)
	if err != nil {
		return err, http.StatusBadRequest
	}

	err = s.Store.DeleteSnacker(id)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	WriteJSON(w, http.StatusOK, nil)
	return nil, http.StatusOK
}

func (s *Server) heartHandler(w http.ResponseWriter, r *http.Request) (error, int) {
	id, err := parseURLID(r)
	if err != nil {
		return err, http.StatusBadRequest
	}

	snackerInit, err := s.Store.FindSnacker(id)
	snackerInit.Hearts += 1
	err = s.Store.UpdateSnacker(snackerInit)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	foundSnacker, err := s.Store.FindSnacker(snackerInit.ID)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	WriteJSON(w, http.StatusOK, foundSnacker)
	return nil, http.StatusOK
}

func parseURLID(r *http.Request) (uint64, error) {
	rawID, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return 0, err
	}
	return uint64(rawID), nil
}
