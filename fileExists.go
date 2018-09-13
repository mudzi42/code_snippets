package codesnippets

//Check if a file exists before using it

import (
	"fmt"
	"os"
)

func exists(fname string) bool {
	if _, err := os.Stat(fname); os.IsNotExist(err) {
		fmt.Printf("File %s does not exist", fname)
		return false
	}
	fmt.Printf("File %s exists", fname)

	return true
}
