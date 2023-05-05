package storage

import (
	"fmt"

	"github.com/hramcovdv/cm_info/models"
)

var headends = map[string]string{
	"10.10.1.11": "public",
	"10.10.1.12": "public",
	"10.10.1.13": "public",
	"10.10.1.14": "public",
	"10.10.1.15": "public",
}

type MemoryStorage struct{}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (m *MemoryStorage) GetHeadend(h *models.Headend, addr string) error {
	if val, ok := headends[addr]; ok {
		h.Addr = addr
		h.Comm = val

		return nil
	}

	return fmt.Errorf("headend not found with IP %s", addr)
}

func (m *MemoryStorage) GetHeadends() ([]models.Headend, error) {
	var a []models.Headend
	for k, v := range headends {
		h := models.Headend{
			Addr: k,
			Comm: v,
		}

		a = append(a, h)
	}

	return a, nil
}
