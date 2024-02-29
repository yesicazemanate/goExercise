package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Panda struct {
	ID          string `json:"id"`
	Nombre      string `json:"nombre"`
	Description string `json:"description"`
}

var animal = []Panda{
	{ID: "1", Nombre: "oso", Description: "es grande "},
	{ID: "2", Nombre: "perro", Description: "son bravos"},
	{ID: "3", Nombre: "gato", Description: "tienen grandes uñas"},
}

func getAlimento(a *gin.Context) {
	a.IndentedJSON(http.StatusOK, animal)

}
func getAnimalByID(c *gin.Context) {

	id := c.Param("id")
	for _, a := range animal {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
}
func main() {
	router := gin.Default()

	router.GET("/alimentos", getAlimento)
	router.GET("/animal/:id", getAnimalByID)
	router.Run("localhost:4000")
}
