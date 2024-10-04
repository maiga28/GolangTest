package data

import (
	"encoding/json"
	"log"
)

// Définir la structure pour les données
type Person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

// Données JSON
var data = `{
	"data": [
		{
			"id": "1",
			"name": "John Doe",
			"age": "27",
			"city": "New York"
		},
		{
			"id": "2",	
			"name": "Jane Doe",	
			"age": "25",	
			"city": "San Francisco"
		},
		{
			"id": "3",	
			"name": "Bob Smith",	
			"age": "32",	
			"city": "Los Angeles"
		},
		{
			"id": "4",	
			"name": "Alice Johnson",	
			"age": "28",	
			"city": "Chicago"
		},
		{
			"id": "5",	
			"name": "Mike Williams",	
			"age": "30",	
			"city": "Seattle"
		}
	]
}`

// Fonction pour décoder les données JSON
func GetPeople() []Person {
	var people struct {
		Data []Person `json:"data"`
	}

	err := json.Unmarshal([]byte(data), &people)
	if err != nil {
		log.Fatal(err)
	}

	return people.Data
}
