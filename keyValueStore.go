package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// data storage map/ hash table
// key is string and value is also string
var dataStore = make(map[string]string)

//add new value to store
func addValue(key string, value string) bool {
	if key == "" {
		return false
	}

	//check if key already exists
	if !existsKey(key) {
		dataStore[key] = value
		return true
	}

	return false
}

// check if key exists in store
func existsKey(key string) bool {
	_, ok := dataStore[key]

	if ok {
		return true
	}

	return false
}

// get value by key
func getValue(key string) string {
	if existsKey(key) {
		value := dataStore[key]
		return value
	}

	return ""
}

// change/update value
func updateValue(key string, value string) bool {
	if existsKey(key) {
		dataStore[key] = value
		return true
	}

	return false
}

// delete value
func deleteValue(key string) bool {
	if existsKey(key) {
		delete(dataStore, key)
		return true
	}

	return false
}

// print all key vlaues
func printValues() {

	if len(dataStore) == 0 {
		message("!!!   Store is empty. add some value first   !!!", false)
	}

	for key, value := range dataStore {
		fmt.Printf("Key: %s \t Value: %s\n", key, value)
	}
}

// parse command
func parseCommand(tokens []string) (string, string) {
	//extact key. key must be in 1 index
	key := tokens[1]
	//extact value.
	value := ""

	if len(tokens) > 2 {
		// Value must be from 2 index
		valueSlice := tokens[2:]
		value = strings.Join(valueSlice, " ")
	}

	return key, value
}

// message
func message(text string, showGuide bool) {
	fmt.Println(text)
	if showGuide {
		usage()
	}

}

//welcome
func welcome() {
	fmt.Println("============= SKVS ===============")
	fmt.Println("| Welcome to Key Value Store     |")
	fmt.Println("| This store has 4 functions     |")
	fmt.Println("| ADD, GET, UPDATE, DELETE value |")
	fmt.Println("==================================")
}

//usage guide
func usage() {
	fmt.Println("============= USAGE ==============")
	fmt.Println("|  add key value                 |")
	fmt.Println("|  get key                       |")
	fmt.Println("|  update key value              |")
	fmt.Println("|  delete key                    |")
	fmt.Println("|  type print to show all values |")
	fmt.Println("|  type exit to stop the program |")
	fmt.Println("|  i.e: add go Go is a lang      |")
	fmt.Println("==================================")
	fmt.Print("> ")
}

//main function
func main() {
	welcome()
	usage()
	// init the scanner
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)

		//check input length
		if len(text) == 0 {
			message("!!! Invalid command   !!!", true)
			continue
		}
		tokens := strings.Fields(text)
		switch strings.ToUpper(tokens[0]) {
		case "EXIT":
			message("Thank you for using this program. :)", false)
			return
		case "PRINT":
			printValues()
			fmt.Print("> ")
		case "ADD":
			//validate command
			if len(tokens) < 3 {
				message("!!!   Invalid command   !!!", true)
				continue
			}
			// parse the command
			key, value := parseCommand(tokens)
			if !addValue(key, value) {
				message("!!!   Add operation failed   !!!", false)
			}
			fmt.Print("> ")

		case "GET":
			//validate command
			if len(tokens) < 2 {
				message("!!!   Invalid command   !!!", true)
				continue
			}
			// parse the command
			key, _ := parseCommand(tokens)
			value := getValue(key)
			if value == "" {
				message("!!!   Key doesn't exists   !!!", false)
			} else {
				fmt.Printf("Value: %s\n", value)
			}
			fmt.Print("> ")

		case "UPDATE":
			//validate command
			if len(tokens) < 3 {
				message("!!!   Invalid command   !!!", true)
				continue
			}
			// parse the command
			key, value := parseCommand(tokens)
			if !updateValue(key, value) {
				message("!!!   Update operation failed   !!!", false)
			}
			fmt.Print("> ")

		case "DELETE":
			//validate command
			if len(tokens) < 2 {
				message("!!!   Invalid command   !!!", true)
				continue
			}
			// parse the command
			key, _ := parseCommand(tokens)
			if !deleteValue(key) {
				message("!!!   Delete operation failed   !!!", false)
			}
			fmt.Print("> ")

		default:
			message("!!!   Unknown command   !!!", true)
		}
	}

}
