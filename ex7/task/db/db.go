package db

import (
	"encoding/binary"
	"fmt"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

const taskList = "Task List"

type Task struct {
	Key   int
	Value string
}

func Initialize(dbName string) error {
	var err error
	db, err = bolt.Open(dbName, 0600, nil)
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists([]byte(taskList))
		return err
	})
}

func Close() {
	db.Close()
}
func Add(task string) error {
	var err error
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(taskList))
		id, _ := b.NextSequence()
		taskID := itob(int(id))
		err = b.Put(taskID, []byte(task))
		return err
	})
	return err
}

func List() string {
	result := "TODO:\n"
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(taskList))

		b.ForEach(func(k, v []byte) error {
			result += fmt.Sprintf("%v   %s\n", btoi(k), v)
			return nil
		})
		return nil
	})
	return result
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(id []byte) int {
	return int(binary.BigEndian.Uint64(id))
}
