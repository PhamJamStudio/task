package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

var cmdCompleted = &cobra.Command{
	Use:   "completed",
	Short: "List all of your completed tasks",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You have the following completed tasks:")
		db, err := bolt.Open(dbFile, 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		err = db.View(func(tx *bolt.Tx) error {
			// Assume bucket exists and has keys
			b := tx.Bucket([]byte(todoBucket))

			c := b.Cursor()
			for k, v := c.First(); k != nil; k, v = c.Next() {
				var t task
				if err := json.Unmarshal(v, &t); err != nil {
					return err
				}
				// only print tasks with status = "to do"
				if t.Status == "done" {
					fmt.Printf("%v. %s\n", btoi(k), t.Name)
				}
			}
			return nil
		})

		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(cmdCompleted)
}
