package utils

import (
	"github.com/gofrs/uuid"
)

func NewUUID() string {
	u1, _ := uuid.NewV4()
	return u1.String()
}
