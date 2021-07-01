package internal

import (
	"log"

	"github.com/boltdb/bolt"
)

func HandleDbConnection() (*bolt.DB, map[string]string) {
	db, err := bolt.Open("urlshort.db", 0600, nil)
	dbMap := make(map[string]string)

	if err != nil {
		log.Fatal(err)
	}

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("pairs"))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			dbMap[string(k)] = string(v)
		}
		return nil
	})

	return db, dbMap
}
