package auth

import (
	"encoding/base64"
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type AuthMiddleware struct {
	realm string
	auth  Authenticator
}

func NewAuthMiddleware(realm string, auth Authenticator) *AuthMiddleware {
	return &AuthMiddleware{
		realm: realm,
		auth:  auth,
	}
}

func (m *AuthMiddleware) Authenticate(c *gin.Context) {
	authHeader := m.extractAuthHeader(c)
	user, pass := m.userNamePair(authHeader)

	authenticated, u := m.auth(user, pass)
	if authenticated {
		c.Set("user", u)
		return
	}
	m.denyAccess(c)
}

func (m *AuthMiddleware) userNamePair(authHeader string) (string, string) {
	if authHeader == "" {
		return "", ""
	}
	decodeString, err := base64.StdEncoding.DecodeString(authHeader)
	if err != nil {
		log.Error(err)
		return "", ""
	}

	userPass := strings.Split(string(decodeString), ":")
	return userPass[0], userPass[1]
}

func (m *AuthMiddleware) extractAuthHeader(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}
	split := strings.Split(authHeader, " ")
	return split[1]
}

func (m *AuthMiddleware) denyAccess(c *gin.Context) {
	authHeader := "Basic realm=" + strconv.Quote(m.realm)
	c.Header("WWW-Authenticate", authHeader)
	c.AbortWithStatus(http.StatusUnauthorized)
}
