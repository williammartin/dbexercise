package dbexercise

import "errors"

type Transaction map[string]string

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		depth:        0,
		transactions: []Transaction{Transaction{}},
	}
}

type InMemoryDB struct {
	depth        int
	transactions []Transaction
}

func (db *InMemoryDB) Set(key string, value string) {
	db.transactions[db.depth][key] = value
}

func (db *InMemoryDB) Get(key string) (string, bool) {
	val, ok := db.transactions[db.depth][key]
	return val, ok
}

func (db *InMemoryDB) BeginTransaction() {
	newTransaction := map[string]string{}
	for k, v := range db.transactions[db.depth] {
		newTransaction[k] = v
	}
	db.transactions = append(db.transactions, newTransaction)
	db.depth = db.depth + 1
}

func (db *InMemoryDB) Commit() error {
	if db.depth == 0 {
		return errors.New("cannot commit without transaction")
	}

	return nil
}

func (db *InMemoryDB) Rollback() error {
	if db.depth == 0 {
		return errors.New("cannot rollback without transaction")
	}

	db.transactions = db.transactions[0:db.depth]
	db.depth = db.depth - 1

	return nil
}

func NewConcurrentInMemoryDB() *ConcurrentInMemoryDB {
	return &ConcurrentInMemoryDB{
		inMemoryDB: NewInMemoryDB(),
	}
}

type ConcurrentInMemoryDB struct {
	inMemoryDB *InMemoryDB
}

func (db *ConcurrentInMemoryDB) Set(key string, value string) {
	db.inMemoryDB.Set(key, value)
}

func (db *ConcurrentInMemoryDB) Get(key string) (string, bool) {
	return db.inMemoryDB.Get(key)
}
