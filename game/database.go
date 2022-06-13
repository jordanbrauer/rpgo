package game

import "github.com/boltdb/bolt"

var storage *bolt.DB

func Storage() *bolt.DB {
	if storage != nil {
		return storage
	}

	var err error

	storage, err = bolt.Open(config.Database.DSN, 0666, nil)

	if err != nil {
		panic(err)
	}

	return storage
}
