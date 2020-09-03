package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

var cmdDo = &cobra.Command{
	Use:   "do",
	Short: "Mark tasks on your TODO list as complete",
	Long:  "Mark one or more tasks on your TODO list as complete e.g task do 1 2 3",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db, err := bolt.Open(dbFile, 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		err = db.Update(func(tx *bolt.Tx) error {
			// Assume bucket exists and has keys
			b := tx.Bucket([]byte(todoBucket))

			/* Loop thru args
			// if not a num, print error msg, continue to next arg
			// if a num
				// if num doesn't exist, print error msg, continue to next arg
				// if num exists, mark status as done
			*/
			for _, key := range args {
				keyInt, err := strconv.Atoi(key)
				// return error if not an int
				if err != nil {
					fmt.Printf("ERROR: \"%v\" is not an number, choose an number representing your task, skipping\n", key)
					continue
				}
				// Get task and update status to done
				keyByte := itob(keyInt)
				v := b.Get([]byte(keyByte))
				if v == nil {
					fmt.Printf("ERROR: \"%v\" doesn't exist as a task, skipping.\n", key)
					continue
				}
				var t task
				if err := json.Unmarshal(v, &t); err != nil {
					return fmt.Errorf("Unmarshal error: %v", err)
				}
				t.Status = "done"

				// Marshal and save updated task to db
				if buf, err := json.Marshal(t); err != nil {
					return err
				} else if err := b.Put([]byte(keyByte), buf); err != nil {
					return fmt.Errorf("Marshal error: %v", err)
				}
				fmt.Printf("You have completed the \"%v\" task.\n", t.Name)
			}

			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(cmdDo)
}
