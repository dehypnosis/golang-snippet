package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/webgenie/go-in-action/chapter2/sample/matchers"
	_ "github.com/webgenie/go-in-action/chapter2/sample/search"
	_ "github.com/webgenie/go-in-action/chapter3/dbdriver/postgres"
)

func init() {
	// standard logger
	log.SetOutput(os.Stdout)
}

func main() {
	log.Println("test..?ss", sql.Drivers()[0])
	if 1 == 2 {
	}
}

// FunctionOfMain blabalbababablba.
func FunctionOfMain() {

}
