package middleware

import (
	"github.com/FAT/repository"
)

func (auth *AuthenticationCtx) CreateToken(account repository.Account) (string, error) {
	payload, err := NewPayload(account)
	if err != nil {
		return "", err
	}

	return auth.paseto.Encrypt(auth.symmetricKey, &payload, payload)
}

func (auth *AuthenticationCtx) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := auth.paseto.Decrypt(token, auth.symmetricKey, payload, nil)
	if err != nil {
		return nil, err
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
