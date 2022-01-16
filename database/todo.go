package database

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/boltdb/bolt"
)

type Task struct {
	Id   int
	Name string
}

const PATH = "/my.db"
const DEFAULT_BUCKET_NAME = "tasks"

func init() {
	db, _ := getDatabase(PATH)
	defer db.Close()
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(DEFAULT_BUCKET_NAME))
		if err != nil {
			return err
		}
		return nil
	})

}

func getDatabase(path string) (*bolt.DB, error) {
	database, err := bolt.Open("../my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return database, nil
}

func AddTask(taskName string) (Task, error) {
	db, _ := getDatabase(PATH)
	todo_id := 0
	defer db.Close()
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(DEFAULT_BUCKET_NAME))
		id, _ := b.NextSequence()
		todo_id = int(id)
		return b.Put([]byte(strconv.Itoa(todo_id)), []byte(taskName))
	})

	if err != nil {
		return Task{}, err
	}

	return Task{Id: todo_id, Name: taskName}, nil
}

func GetAllTaskItems() ([]Task, error) {
	tasks := []Task{}

	db, _ := getDatabase(PATH)
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(DEFAULT_BUCKET_NAME))

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			id, err := strconv.Atoi(string(k))
			if err != nil {
				return err
			}
			tasks = append(tasks, Task{Id: id, Name: string(v)})
		}
		return nil
	})

	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func RemoveTaskItem(taskId int) error {

	db, _ := getDatabase(PATH)
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(DEFAULT_BUCKET_NAME))
		c := b.Cursor()

		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			id, err := strconv.Atoi(string(k))
			if err != nil {
				return err
			}
			if id == taskId {
				c.Delete()
				return nil
			}
		}
		return errors.New("no task with specified id found")
	})

	if err != nil {
		return err
	}
	return nil
}
