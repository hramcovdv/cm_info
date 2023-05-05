package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/hramcovdv/cm_info/models"
	"github.com/hramcovdv/cm_info/storage"
)

type Server struct {
	store storage.Storer
}

func NewServer(store storage.Storer) *Server {
	return &Server{
		store: store,
	}
}

func (s *Server) Start(addr string) error {
	http.HandleFunc("/cm_info", s.getCmIndex)
	http.HandleFunc("/cmts_info", s.getHeadend)
	http.HandleFunc("/cmts_list", s.getHeadends)

	return http.ListenAndServe(addr, nil)
}

func (s *Server) getHeadend(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	addr := r.FormValue("cmts")

	var headend models.Headend
	if err := s.store.GetHeadend(&headend, addr); err != nil {
		log.Printf("GetHeadend() error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	s.WriteJSON(w, headend)
}

func (s *Server) getHeadends(w http.ResponseWriter, r *http.Request) {
	headends, err := s.store.GetHeadends()
	if err != nil {
		log.Printf("GetHeadends() error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	s.WriteJSON(w, headends)
}

func (s *Server) getCmIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	mac := r.FormValue("mac")
	if mac == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	headends, err := s.store.GetHeadends()
	if err != nil {
		log.Printf("GetHeadends() error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var wg sync.WaitGroup

	found := make(chan string, len(headends))
	for _, h := range headends {
		wg.Add(1)

		go func(h models.Headend) {
			defer wg.Done()

			h.GetModemOid(mac, found)
		}(h)
	}

	wg.Wait()
	close(found)

	for f := range found {
		fmt.Print(f)
	}
}

func (s *Server) WriteJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(v)
}
