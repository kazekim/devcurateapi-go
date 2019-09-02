/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package usecase

import (
	"github.com/kazekim/devcurateapi-go/api/app/domain/model"
	"github.com/kazekim/devcurateapi-go/api/app/domain/service"
	"github.com/kazekim/devcurateapi-go/api/app/interface/mongodb"
	"gopkg.in/mgo.v2"
)



type keyUsecase struct {
	service *service.KeyService
}

func NewKeyUsecase(service *service.KeyService) *keyUsecase {
	return &keyUsecase{
		service: service,
	}
}

func BuildKeyUseCase(db *mgo.Database) (*keyUsecase, error) {

	r := mongodb.NewKeyRepository(db)
	s := service.NewKeyService(r)
	u := NewKeyUsecase(s)

	return u, nil
}

func (u *keyUsecase) CreateKey(value string) (*model.Key, error) {

	key, err := u.service.CreateKey(value)
	return key, err
}

func (u *keyUsecase) RemoveKeyCreatedMoreThanOnHour() error {
	return u.service.RemoveKeyCreatedMoreThanOnHour()
}

func (u *keyUsecase) GetKeyByID(id string) (*model.Key, error) {
	key, err := u.service.GetKeyByID(id)
	return key, err
}

func (u *keyUsecase) CheckValidForKey(value string) error {
	err := u.service.CheckValidForKey(value)
	return err
}