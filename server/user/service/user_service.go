package service

import (
	"context"
	"ecs-test/server/user/domain"
	"ecs-test/server/user/repository"
	"github.com/charmbracelet/log"
)

type Service struct {
	userRepository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *Service {
	return &Service{
		userRepository: repository,
	}
}

func (u Service) FindUser(username string, ctx context.Context) *domain.User {
	//ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	//defer cancel()

	userEntity, err := u.userRepository.FetchByUsername(ctx, username)
	if err != nil {
		log.Print(err)
		return nil
	}

	if userEntity == nil {
		return nil
	}
	return userEntity.ToUser()
}

func (u Service) FindById(id uint, ctx context.Context) *domain.User {
	//ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	//defer cancel()

	userEntity, err := u.userRepository.FetchById(ctx, id)
	if err != nil {
		log.Error(err)
		return nil
	}

	if userEntity == nil {
		return nil
	}
	return userEntity.ToUser()
}
