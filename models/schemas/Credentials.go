package schemas

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Credentials struct {
	Id        uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Roles_id  uuid.UUID      `json:"roles_id" gorm:"type:uuid"`
	Roles     *Roles         `json:"roles" gorm:"foreignKey:Roles_id;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Username  string         `json:"username" gorm:"type:varchar(50)"`
	Password  string         `json:"password" gorm:"type:varchar(255)"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
