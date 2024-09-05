package service

import (
	"crypto/sha1"
	"fmt"
	restApi "github.com/RINcHIlol/rest.git"
	"github.com/RINcHIlol/rest.git/pkg/repository"
)

const salt = "hewybq3q89gq9"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user restApi.User) (int, error) {
	user.Password = generatePasswordhash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordhash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
