package rest

import (
	"ecs-test/server/auth"
	"ecs-test/server/infra/http"
	"ecs-test/server/session/container"
	"ecs-test/server/user/domain"
	"ecs-test/server/user/service"
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

type SessionController struct {
	authService      *service.AuthenticationService
	sessionContainer *container.SessionContainer
}

func NewSessionController(
	authService *service.AuthenticationService,
	sessionContainer *container.SessionContainer,
) *SessionController {
	return &SessionController{
		authService:      authService,
		sessionContainer: sessionContainer,
	}
}

func (t *SessionController) RegisterRoutes(r http.RouteProvider) {
	authMiddleware := auth.NewAuthMiddleware("Auth", t.authService.Authenticate)

	r.Route().
		Group("/session", authMiddleware.Authenticate).
		POST("/", t.createSession)
}

func (t *SessionController) createSession(c *gin.Context) {
	value, exists := c.Get("user")
	if exists {
		user, ok := value.(*domain.User)
		if ok {
			existingSession, found := t.sessionContainer.FindExistingSessionForUser(user.Name)
			if found {
				c.Header("X-Session-ID", existingSession.SessionId.String())
				c.Status(200)
				return
			}

			session, err := t.sessionContainer.NewSessionForUser(user.Name)
			if err != nil {
				log.Error(err)
				c.Status(500)
				return
			}
			c.Header("X-Session-ID", session.String())
			c.Status(200)
			return
		}
	}
	c.Status(404)
}
