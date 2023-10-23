package repositories

import "villa_go/entities/models"

type VillaRepository interface {
	GetAllVilla() ([]models.VillaListResponse, error)
	CreateVilla()
	DeleteVilla() (bool, error)
	UpdateVilla()
}
