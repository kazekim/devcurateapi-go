/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package repository

import "github.com/kazekim/devcurateapi-go/api/app/domain/entity"

type KeyRepository interface {
	GetKeyByID(id string) (*entity.Key, error)
	CreateKey(value string) (*entity.Key, error)
	RemoveKey(id string) error
	RemoveKeyTimeLessThan(time int64) error
	CheckValidForKey(value string, time int64) error
}
