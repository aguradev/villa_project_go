package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Admin struct {
	Id            uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Credential_id uuid.UUID      `json:"credential_id" gorm:"type:uuid;"`
	Credential    *Credentials   `json:"credential" gorm:"foreignKey:Credential_id;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	First_name    string         `json:"first_name" gorm:"type:varchar(50)"`
	Last_name     string         `json:"last_name" gorm:"type:varchar(50)"`
	CreatedAt     time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"autoUpdateTime:milli"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
