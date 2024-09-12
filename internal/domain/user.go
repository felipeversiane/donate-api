package domain

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/google/uuid"
)

type UserDomain struct {
	ID         uuid.UUID
	Name       string
	Email      string
	Password   string
	Phone      string
	Document   string
	PostalCode string
	CoverID    uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

func NewUserDomain(
	name,
	email,
	password,
	phone,
	document,
	postalCode string,
	coverID uuid.UUID) *UserDomain {
	return &UserDomain{
		ID:         uuid.New(),
		Name:       name,
		Email:      email,
		Password:   password,
		Phone:      phone,
		Document:   document,
		PostalCode: postalCode,
		CoverID:    coverID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

func (u *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(u.Password))
	u.Password = hex.EncodeToString(hash.Sum(nil))
}
