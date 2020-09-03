/* REQUIREMENTS
//  CLI tool that can be used to manage your TODOs in the terminal

// TASK STRUCTURE
	// ID
	// name
	// status: done, to do
	// Completion date

// SUB COMMANDS
	// add         Add a new task to your TODO list
	// do          Mark a task on your TODO list as complete
	// list        List all of your incomplete tasks
	// BONUS
		// delete	   Delete a task from your TODO list
		// completed   lists tasks finished today
		//
*/
package main

import (
	"github.com/PhamJamStudio/task/cmd"
)

func main() {
	cmd.Execute()

}

/* LEARNING NOTES
// encoding/binary simple translation between num/byte sequences and encoding and decoding of varints.
	func itob(v int) []byte {
		b := make([]byte, 8)
		// BigEndian by having high order byte come first,  allows db's to efficiently determine if # is pos/neg. #'s also store in order they are printed, so binary to decimal routines are very efficient
		binary.BigEndian.PutUint64(b, uint64(v))
		return b
	}
	func btoi(b []byte) int {
		return int(binary.BigEndian.Uint64(b))
	}
// JSON marshall/unmarshall requires structs that make its vals public
// BoltDB requires storing keys in big endian []bytes
// Cobra: CLI library
	// cobra init github.com/gophercises/taskauto -> generates stub files for cobra CLI app
	// cobra add list.go -> creates stub for command incl. rootCmd.
// BoltDB = key/val store, all in go, so no ext deps, single file based
	// bucket ~= tables, key/val pairs in []byte format
	// great for high read, slow write, concurrent writes

// POTENTIAL REFACTOR
	//  Create funcs for each command e.g CreateTask(task string) (int, error)
	// BoltDB
		// Take out errors because in tx, errors only occur if tx is closer or not writable tx
		// use go-homedir to ensure db is always in same location + filepath.Join(home, "tasksy .db")
		// Enable user submitted db path
		// put db in own db folder
		// Add e.g &bolt.Options{Timeout: 1 * time.Second}) since bolt.Open locks
		// Redesign w/ nested  buckets vs current json design or storm pkg for boltDB e.g
		UserBucket
		key		|	value
		id123	|	some bucket

		SomeBucket
		key		|	value
		name	|	"Yen"
		Email	|	...
	// add must func - https://courses.calhoun.io/lessons/les_goph_39
	// Don't use package lvl vars for real web apps, use structs?


*/

/*
// ENHANCEMENT IDEAS
	// Handling unicode, concurrent users? Write your own DB layer, replace w/ own tooling

// DESIGN
// Build CLI shell task command and ssub commands above stubs
	// add
	// do
	// list
// Write BoltDB interactions
	// add
		// check if task exists, add if it doesn't
	// do
		// check if task exists, mark as done if so
// Enable app install to run from anywhere and persist tasks
	// Look into finding user home dir on any OS e.g https://github.com/mitchellh/go-homedir
	// installing binary -> go install --help, OS specific issues https://github.com/gophercises/task/issues?utf8=%E2%9C%93&q=is%3Aissue
*/
