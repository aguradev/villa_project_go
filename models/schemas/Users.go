package schemas

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Users struct {
	Id              *uuid.UUID     `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Credential_id   uuid.UUID      `json:"credential_id" gorm:"type:uuid;"`
	Credential      *Credentials   `json:"credential" gorm:"foreignKey:Credential_id;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	First_name      string         `json:"first_name,omitempty" gorm:"type:varchar(50)"`
	Last_name       string         `json:"last_name,omitempty" gorm:"type:varchar(50)"`
	Email           string         `json:"email,omitempty" gorm:"type:varchar(50);unique"`
	Address         string         `json:"address,omitempty" gorm:"type:varchar(255)"`
	Profile_picture string         `json:"profile_picture,omitempty" gorm:"type:varchar(255)"`
	CreatedAt       time.Time      `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `json:"updated_at,omitempty" gorm:"autoUpdateTime:milli"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
