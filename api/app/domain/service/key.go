/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package service

import (
	"github.com/kazekim/devcurateapi-go/api/app/domain/model"
	"github.com/kazekim/devcurateapi-go/api/app/domain/repository"
	"github.com/mitchellh/mapstructure"
	"time"
)

type KeyService struct {
	repo repository.KeyRepository
}

func NewKeyService(repo repository.KeyRepository) *KeyService {
	return &KeyService{
		repo: repo,
	}
}

func (s *KeyService) CreateKey(value string) (*model.Key, error) {

	key, err := s.repo.CreateKey(value)
	if err != nil {
		return nil, err
	}

	var m model.Key
	err = mapstructure.Decode(key, &m)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (s *KeyService) RemoveKeyCreatedMoreThanOnHour() error {

	deleteTime := time.Now().Unix() - 3600

	err := s.repo.RemoveKeyTimeLessThan(deleteTime)
	if err != nil {
		return err
	}
	return nil
}

func (s *KeyService) GetKeyByID(id string) (*model.Key, error) {

	key, err := s.repo.GetKeyByID(id)
	if err != nil {
		return nil, err
	}

	var m model.Key
	err = mapstructure.Decode(key, &m)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (s *KeyService) CheckValidForKey(value string) error {

	deleteTime := time.Now().Unix() - 3600
	err := s.repo.CheckValidForKey(value, deleteTime)
	if err != nil {
		return err
	}

	return nil
}