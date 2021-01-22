package jojivm

//General utility functions

import (
	"log"
	"os"
)

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
