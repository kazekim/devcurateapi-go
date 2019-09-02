/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package mongofw

import (
	"gopkg.in/mgo.v2"
	"log"
)

type Mongo struct {
	session *mgo.Session
}

func Open() (*mgo.Session, error) {
	session, err := mgo.Dial("mongo:27017")
	if err != nil {
		log.Fatalln("mongo err")
		return nil, err
	}


	session.SetMode(mgo.Monotonic, true)

	return session, nil
}

func (m *Mongo) Close() {
	m.session.Close()
}

func (m *Mongo) DB(name string) *mgo.Database {
	return m.session.DB(name)
}