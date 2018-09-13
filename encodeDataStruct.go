package codesnippets

type Address struct {
	Line1    string `json:"line_1"`
	City     string `json:"city"`
	Postcode string `json:"postcode"`
}

// Mutliple Tags

type Person struct {
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
}

// Ignore Empty

type Animal struct {
	Name  string `json:"name"`
	Noise string `json:"noise,omitempty"`
	Size  string `json:",omitempty"`
}
