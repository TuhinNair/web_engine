// +build integration
// +build mgo

package mongo

import (
	"log"
	"os"
	"testing"
	"webengine/data/model"

	mgo "gopkg.in/mgo.v2"
)

const (
	dbName = "mongo"
)

var (
	db *mgo.Session
)

func TestMain(m *testing.M) {
	conn, err := model.Open("mongo", "127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}

	db = conn
	defer conn.Close()
	os.Exit(m.Run())
}
