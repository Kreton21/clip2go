package bolted

import (
	"fmt"

	"github.com/boltdb/bolt"
)

func Wdb(bucket, key, value []byte) {
	db, err := bolt.Open("clip2.db", 0600, nil)
	if err != nil {
		fmt.Print(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		err := b.Put(key, value)
		return err
	})
	db.Close()
}

func Rdb(bucket, key string) string {
	var result string
	db, err := bolt.Open("clip2.db", 0600, nil)
	if err != nil {
		fmt.Print(err)
	}

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		v := b.Get([]byte(key))
		result = string(v)
		//	fmt.Printf(string(v))
		return nil
	})
	db.Close()
	return result
}
