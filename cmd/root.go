package cmd

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

var db *bolt.DB
var dbFile = "todo.db"
var todoBucket = []byte("todos")

type task struct {
	Name     string
	Status   string
	doneTime time.Time
}

func initBoltDB(db *bolt.DB) error {
	// open or create boltDB
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists(todoBucket)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// btoi return an int representation of a byte big endian
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "task is a CLI for managing your TODOs.",
}

// Execute executes the root command.
func Execute() {
	err := initBoltDB(db)
	if err != nil {
		log.Fatal(err)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
