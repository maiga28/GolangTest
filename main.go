package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/maiga28/golangtest/data" // Importez le package avec le bon chemin
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		// Appel de la fonction GetPeople
		people := data.GetPeople()

		// Affichage dans la console
		for _, person := range people {
			fmt.Printf("ID: %s, Name: %s, Age: %d, City: %s\n", person.ID, person.Name, person.Age, person.City)
		}

		// Envoi de la réponse JSON avec la liste des personnes
		c.JSON(200, gin.H{
			"people": people,
		})
	})

	// Lancement du serveur sur 0.0.0.0:8080
	r.Run() // listen and serve on 0.0.0.0:8080
}
