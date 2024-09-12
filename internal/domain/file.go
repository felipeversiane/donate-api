package domain

import (
	"github.com/google/uuid"
)

type FileDomain struct {
	ID  uuid.UUID
	Key string
	URL string
}

func NewFileDomain(url, key string) FileDomain {
	return FileDomain{
		ID:  uuid.New(),
		Key: key,
		URL: url,
	}
}
