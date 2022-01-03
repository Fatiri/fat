package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/fat/common/wrapper"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

type Role string

// List of role
var (
	RoleAdmin Role = "ADMIN"
	RoleUser  Role = "User"
)

// AuthMiddleware creates a gin middleware for authorization
func (maker *MakerCtx) AuthMiddleware(roles []Role) gin.HandlerFunc {
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
		payload, err := maker.VerifyToken(accessToken)
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
