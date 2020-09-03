package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

var cmdAdd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// open db
		db, err := bolt.Open(dbFile, 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		// Access bucket
		err = db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket(todoBucket)

			// Generate key for new task
			id64, _ := b.NextSequence()
			key := itob(int(id64))
			t := task{
				Name:   strings.Join(args, " "),
				Status: "to do",
			}

			// Marshal and save encoded task
			if buf, err := json.Marshal(t); err != nil {
				return err
			} else if err := b.Put(key, buf); err != nil {

				return err
			}

			fmt.Printf("Added task \"%v\" to your task list.", t.Name)
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(cmdAdd)
}
