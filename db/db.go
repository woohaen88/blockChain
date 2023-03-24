package db

import (
	"github.com/boltdb/bolt"
	"github.com/woohaen88/utils"
)


var db *bolt.DB

const (
	dbName = "blockchain.db"
	dataBucket = "data"
	blocksBucket = "blocks"
)

func DB() *bolt.DB{
	if db == nil {
		 dbPointer, err := bolt.Open(dbName, 0600, nil)
		 db = dbPointer
		 utils.HandleErr(err)
		 err = db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte(dataBucket))
			utils.HandleErr(err)

			_, err = tx.CreateBucketIfNotExists([]byte(blocksBucket))
			utils.HandleErr(err)

			return err
		 })
		 
	}
	return db
}