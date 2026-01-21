package domain

import "github.com/google/uuid"

type UUID string

func NewUUID() UUID {
	return UUID(uuid.NewString())
}
