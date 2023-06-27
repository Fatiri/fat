package middleware

import (
	"errors"
	"time"

	"github.com/Fatiri/fat/repository"
	"github.com/google/uuid"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Type      string    `json:"type"`
	IssuedAt  int64     `json:"issued_at"`
	ExpiredAt int64     `json:"expired_at"`
}

func NewPayload(account repository.Account) (*Payload, error) {
	var duration time.Duration
	if account.AccountType == string(RoleUser) {
		duration = 24
	} else {
		duration = 168
	}

	timeAsiaJakarta, _ := time.LoadLocation("Asia/Jakarta")
	start := time.Now().In(timeAsiaJakarta).UTC()
	end := start.Add(time.Hour * duration)

	payload := &Payload{
		ID:        account.AccountID,
		Username:  account.Username,
		Type:      account.AccountType,
		IssuedAt:  start.Unix(),
		ExpiredAt: end.Unix(),
	}

	return payload, nil
}

var ErrExpiredToken = errors.New("token has expired")

func (payload *Payload) Valid() error {
	exp := time.Unix(payload.ExpiredAt, 0)
	if time.Now().After(exp) {
		return ErrExpiredToken
	}
	return nil
}
