package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

const (
	DonationTypeOneTime   = "one_time"
	DonationTypeRecurring = "recurring"
)

const (
	PaymentMethodDebitCard  = "debit_card"
	PaymentMethodCreditCard = "credit_card"
	PaymentMethodPix        = "pix"
)

type DonateDomain struct {
	ID                uuid.UUID
	Amount            float64
	UserID            uuid.UUID
	PaymentMethod     string
	InstallmentNumber int
	DonationType      string
	Period            int
	CreatedAt         time.Time
}

func NewDonateDomain(amount float64, userID uuid.UUID, paymentMethod, donationType string, installmentNumber, period int) *DonateDomain {
	return &DonateDomain{
		ID:                uuid.New(),
		Amount:            amount,
		UserID:            userID,
		PaymentMethod:     paymentMethod,
		InstallmentNumber: installmentNumber,
		DonationType:      donationType,
		Period:            period,
		CreatedAt:         time.Now(),
	}
}

func (d *DonateDomain) Validate() error {
	if d.DonationType != DonationTypeOneTime && d.DonationType != DonationTypeRecurring {
		return errors.New("invalid donation type")
	}
	if d.PaymentMethod != PaymentMethodDebitCard && d.PaymentMethod != PaymentMethodCreditCard && d.PaymentMethod != PaymentMethodPix {
		return errors.New("invalid payment method")
	}
	if d.DonationType == DonationTypeOneTime && d.PaymentMethod == PaymentMethodCreditCard && d.InstallmentNumber != 0 {
		return errors.New("one-time donations with credit card must have installment number as 0")
	}
	if d.DonationType == DonationTypeRecurring && d.PaymentMethod != PaymentMethodCreditCard {
		return errors.New("recurring donations can only be made with credit card")
	}
	if d.DonationType == DonationTypeRecurring && d.Period <= 0 {
		return errors.New("recurring donations must have a period greater than 0")
	}
	if d.DonationType == DonationTypeOneTime && d.Period != 0 {
		return errors.New("one-time donations must have period as 0")
	}
	return nil
}
