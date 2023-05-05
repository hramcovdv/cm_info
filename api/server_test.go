package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hramcovdv/cm_info/models"
	"github.com/hramcovdv/cm_info/storage"
)

func TestGetHeadend(t *testing.T) {
	db := storage.NewMemoryStorage()
	srv := NewServer(db)

	url := fmt.Sprintf("/cmts_info?cmts=%s", "10.10.1.13")

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(srv.getHeadend)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("unexpected code: want %d got %d",
			http.StatusOK, rr.Code)
	}

	var expected models.Headend
	if err := db.GetHeadend(&expected, "10.10.1.13"); err != nil {
		t.Errorf("GetHeadend() error: %s", err)
	}

	var headend models.Headend
	json.NewDecoder(rr.Body).Decode(&headend)

	if headend != expected {
		t.Errorf("unexpected headend: want %v got %v", headend, expected)
	}
}

func TestGetHeadends(t *testing.T) {
	db := storage.NewMemoryStorage()
	srv := NewServer(db)

	req, _ := http.NewRequest(http.MethodGet, "/cmts_list", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(srv.getHeadends)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("unexpected code: want %d got %d",
			http.StatusOK, rr.Code)
	}

	expected, err := db.GetHeadends()
	if err != nil {
		t.Errorf("GetHeadends() error: %s", err)
	}

	var headends []models.Headend

	json.NewDecoder(rr.Body).Decode(&headends)
	if len(headends) != len(expected) {
		t.Errorf("unexpected headends count: want %v got %v",
			len(headends), len(expected))
	}
}

func TestGetCmIndex(t *testing.T) {
	db := storage.NewMemoryStorage()
	srv := NewServer(db)

	url := fmt.Sprintf("/cm_info?mac=%s", "00:00:bb:ff:cc:dd")

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(srv.getCmIndex)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("unexpected code: want %d got %d",
			http.StatusOK, rr.Code)
	}
}
