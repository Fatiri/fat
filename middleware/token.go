package middleware

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/fat/repository"
	"github.com/gin-gonic/gin"
	"github.com/o1egl/paseto"
)

type Maker interface {
	CreateToken(account repository.Account, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
	AuthMiddleware(roles []Role) gin.HandlerFunc
}

type MakerCtx struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &MakerCtx{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}

func (maker *MakerCtx) CreateToken(account repository.Account, duration time.Duration) (string, error) {
	payload, err := NewPayload(account, duration)
	if err != nil {
		return "", err
	}

	return maker.paseto.Encrypt(maker.symmetricKey, &payload, nil)
}

func (maker *MakerCtx) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, err
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
