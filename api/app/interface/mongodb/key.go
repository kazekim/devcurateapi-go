/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package mongodb

import (
	"github.com/kazekim/devcurateapi-go/api/app/domain/entity"
	"github.com/kazekim/devcurateapi-go/api/pkg/jhrandom"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const KeyCollectionkName = "key"

type keyRepository struct {
	collection *mgo.Collection
}

func NewKeyRepository(db *mgo.Database) *keyRepository {

	c := db.C(KeyCollectionkName)

	return &keyRepository{
		collection: c,
	}
}

func (r *keyRepository) GetKeyByID(id string) (*entity.Key, error) {

	var result entity.Key
	err := r.collection.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *keyRepository) CheckValidForKey(value string, time int64) error {
	var result entity.Key
	err := r.collection.Find(bson.M{"value": value, "createat" : bson.M{"$lt": time}}).One(&result)
	if err != nil {
		return err
	}
	return nil
}

func (r *keyRepository) CreateKey(value string) (*entity.Key, error) {

	id := jhrandom.String(8)
	key := entity.Key{
		ID:id,
		Value: value,
	}
	key.CreatedAt = time.Now().Unix()

	if err := r.collection.Insert(key); err != nil {
		return nil, err
	}

	return &key, nil
}

func (r *keyRepository) RemoveKey(id string) error {
	if err := r.collection.Remove(bson.M{"id": id}); err != nil {
		return err
	}

	return nil
}

func (r *keyRepository) RemoveKeyTimeLessThan(time int64) error {
	_, err := r.collection.RemoveAll(bson.M{"createdat" : bson.M{"$lt": time}})
	if err != nil {
		return err
	}
	return nil
}