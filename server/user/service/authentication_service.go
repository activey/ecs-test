package service

import (
	"context"
	"ecs-test/server/user/domain"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthenticationService struct {
	userService *Service
}

func NewAuthenticationService(userService *Service) *AuthenticationService {
	return &AuthenticationService{
		userService: userService,
	}
}

func (a *AuthenticationService) Authenticate(user string, password string) (bool, *domain.User) {
	authenticated := a.FindAuthenticated(user, password)
	if authenticated != nil {
		return true, authenticated
	}
	return false, nil
}

func (a *AuthenticationService) checkPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil // if err is nil, it means the passwords match
}

func (a *AuthenticationService) FindAuthenticated(username, password string) *domain.User {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	u := a.userService.FindUser(username, ctx)
	if u == nil {
		return nil
	}
	if a.checkPassword(u.Password, password) {
		return u
	}
	return nil
}
