package designpatterns

type DB struct {

}

var db *DB

func initDB() *DB {
	return &DB{}
}

func GetDB() *DB {
	if db == nil {
		db = initDB()
	}

	return db
}