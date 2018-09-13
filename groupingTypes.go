package codesnippets

//Grouping Types

import "fmt"

// Speaker speak[ER] provides a common behavior for all concrete types
// to follow if they want to be part of this group.  This
// is a contract for these concrete types to follow.
type Speaker interface {
	Speak()
}

// Dog contains everything a Dog needs
type Dog struct {
	Name       string
	IsMammal   bool
	PackFactor int
}

// Speak knows how to speak like a dog.
// This makes a Dog now part of a group of concrete
// type that know how to speak.
func (d Dog) Speak() {
	fmt.Println("Woof!",
		"My name is,", d.Name,
		", it is", d.IsMammal,
		"I am a mammal with a pack factor of ", d.PackFactor)
}

// Cat contains everything a Cat needs
type Cat struct {
	Name        string
	IsMammal    bool
	ClimbFactor int
}

func (c Cat) Speak() {
	fmt.Println("Meow!",
		"My name is,", c.Name,
		", it is", c.IsMammal,
		"I am a mammal with a climb factor of ", c.ClimbFactor)
}

func Callmain() {
	fmt.Println("Animal Example:")

	// create a list of Animals that know how to speak
	speakers := []Speaker{
		Dog{
			Name:       "Fido",
			IsMammal:   true,
			PackFactor: 5,
		},
		Cat{
			Name:        "Whiskers",
			IsMammal:    true,
			ClimbFactor: 5,
		},
	}

	// Have the Animals speak
	for _, spkr := range speakers {
		spkr.Speak()
	}

}

//Here are some guidelines around declaring types:
// - Declare types that represent something new or unique
// - Validate that a value of any type is created or used on its own
// - Embed types to reuse existing BEHAVIORS you need to satisfy
// - Question types that are an alias or abstraction for an existing type
// - Question types whose sole purpose is to share common state
