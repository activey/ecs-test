package auth

import "ecs-test/server/user/domain"

type Authenticator func(user, password string) (bool, *domain.User)
