package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userid"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)

	if header == "" {
		NewErrorResponse(c, http.StatusUnauthorized, "Empty authorization header")
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		NewErrorResponse(c, http.StatusUnauthorized, "Invalid authorization header")
		return
	}

	userid, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userid)
}

func getUsersId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "User id did not found in context")
		return 0, errors.New("user id did not found in context")
	}

	idInt, ok := id.(int)

	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "User id is of invalid type")
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
