package request

import uuid "github.com/satori/go.uuid"

type GalleryRequest struct {
	Id      *uuid.UUID
	Fileurl []string
}
