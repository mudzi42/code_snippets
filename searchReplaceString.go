package codesnippets

//Search and Replace in a String

import (
	"fmt"
	"strings"
)

func searchReplaceString() {
	// Example 1: Basic
	myText := "Welcome to GoLangCode.com"
	myText = strings.Replace(myText, "Welcome", "Willkommen", -1)

	// Output: Willkommen to GoLangCode.com
	fmt.Println(myText)

	// Example 2: Change first occurance
	myText = "The sound sounds sound"
	myText = strings.Replace(myText, "sound", "car", 1)

	// Output: The car sounds sound
	fmt.Println(myText)

	// Example 3: Replacing quotes (double backslash needed)
	myText = "I 'quote' this text"
	myText = strings.Replace(myText, "'", "\\'", -1)

	// Output: I \'quote\' this text
	fmt.Println(myText)
}
