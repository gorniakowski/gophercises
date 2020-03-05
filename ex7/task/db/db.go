package db

import (
	"encoding/binary"

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
func Add(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(taskList))
		id64, _ := b.NextSequence()
		taskID := itob(int(id64))
		return b.Put(taskID, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

func List() []Task {
	var result []Task
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(taskList))

		b.ForEach(func(k, v []byte) error {
			task := Task{Key: btoi(k), Value: string(v)}
			result = append(result, task)
			return nil
		})
		return nil
	})
	return result
}

func Remove(id int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(taskList))
		key := itob(id)
		err := b.Delete(key)
		return err
	})
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
