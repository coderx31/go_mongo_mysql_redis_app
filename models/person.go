package models

type Person struct {
	Firstname string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Age       int    `json:"age,omitempty" bson:"age,omitempty"`
	City      string `json:"city,omitempty" bson:"city,omitempty"`
}
