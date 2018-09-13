package codesnippets

import "fmt"

func printStruct(name string) {
	// To print the name of the fields in a struct:

	fmt.Printf("%+v\n", name)
}
