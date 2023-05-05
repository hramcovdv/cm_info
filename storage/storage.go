package storage

import "github.com/hramcovdv/cm_info/models"

type Storer interface {
	GetHeadend(*models.Headend, string) error
	GetHeadends() ([]models.Headend, error)
}
