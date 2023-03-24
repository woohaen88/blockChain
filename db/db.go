package db

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/woohaen88/utils"
)


var db *bolt.DB

const (
	dbName = "blockchain.db"
	dataBucket = "data"
	blocksBucket = "blocks"
	checkpoint = "checkpoint"
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

func SaveBlock(hash string, data []byte) {
	fmt.Printf("Saving Block %s\n", hash)
	err := DB().Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		err := bucket.Put([]byte(hash), data)
		return err
	})
	utils.HandleErr(err)
}

func SaveBlockChain(data []byte) {
	err := DB().Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(dataBucket))
		err := bucket.Put([]byte(checkpoint), data)
		return err
	})

	utils.HandleErr(err)
}

func Checkpoint() []byte {
	// data버킷에서 checkpoint -> []byte로 변환
	var data []byte
	DB().View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(dataBucket))
		data = bucket.Get([]byte(checkpoint))
		return nil
	})
	return data
}

func Block(hash string) []byte{
	var data []byte
	DB().View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		data = bucket.Get([]byte(hash))
		return nil
	})
	return data
}

func Close(){
	DB().Close()
}