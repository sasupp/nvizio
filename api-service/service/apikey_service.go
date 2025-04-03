package service

import (
	"crypto/sha256"
	"crypto/subtle"

	"xtrinio.com/repository"
)

var (
	apiKeyRoot = "projectg"
)

type IApikeyService interface {
	IsApikeyValid(apikey string) bool
}

type apikeyServiceImpl struct {
	repo repository.ApikeyRepository
}

func (a apikeyServiceImpl) IsApikeyValid(apiKey string) bool {
	apiKeyHash := sha256.Sum256([]byte(apiKey))
	isValid, _ := a.repo.IsApikeyHashAvailable(string(apiKeyHash[:]))
	return isValid
}

type localApikeyServiceImpl struct {
	repo repository.ApikeyRepository
}

func (a localApikeyServiceImpl) IsApikeyValid(apiKey string) bool {
	apiKeyHash := sha256.Sum256([]byte(apiKey))
	apiKeyHasRoot := sha256.Sum256([]byte(apiKeyRoot))
	return subtle.ConstantTimeCompare(apiKeyHash[:], apiKeyHasRoot[:]) == 1
}

type allowAllKeyServiceImpl struct {
}

func (a allowAllKeyServiceImpl) IsApikeyValid(apiKey string) bool {
	return true
}

func NewApikeyService(repo repository.ApikeyRepository) IApikeyService {
	return allowAllKeyServiceImpl{}
}
