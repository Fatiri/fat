package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/FAT/common/wrapper"
	"github.com/FAT/repository"
	"github.com/aead/chacha20poly1305"
	"github.com/gin-gonic/gin"
	"github.com/o1egl/paseto"
)

const (
	authorizationHeaderKey  = "Authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

type Role string

// List of role
var (
	RoleAdmin Role = "ADMIN"
	RoleUser  Role = "USER"
)

type Authentication interface {
	CreateToken(account repository.Account) (string, error)
	VerifyToken(token string) (*Payload, error)
	AuthMiddleware(roles []Role) gin.HandlerFunc
}

type AuthenticationCtx struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewAuthentication(symmetricKey string) (Authentication, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	auth := &AuthenticationCtx{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return auth, nil
}

// AuthMiddleware creates a gin middleware for authorization
func (auth *AuthenticationCtx) AuthMiddleware(roles []Role) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, wrapper.ErrorHandler(err))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, wrapper.ErrorHandler(err))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, wrapper.ErrorHandler(err))
			return
		}

		accessToken := fields[1]
		payload, err := auth.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, wrapper.ErrorHandler(err))
			return
		}

		isAuthorized := false
		for _, role := range roles {
			if string(role) == payload.Type {
				isAuthorized = true
			}
		}

		if !isAuthorized {
			err := errors.New("forbiden access")
			ctx.AbortWithStatusJSON(http.StatusForbidden, wrapper.ErrorHandler(err))
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
