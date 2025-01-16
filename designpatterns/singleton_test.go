package designpatterns

import "sync"

type DB struct {
}

var db *DB
var once sync.Once

func initDB() *DB {
	return &DB{}
}

func GetDB() *DB {
	once.Do(func() {
		db = initDB()
	})

	return db
}
