// //Edinson
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Aranass struct {
	ID          string `json:"id"`
	Nombre      string `json:"nombre"`
	Description string `json:"description"`
}

var aranass = []Aranass{
	{ID: "1", Nombre: "Tarantulas", Description: "arañas migalomorfas de gran tamaño con el cuerpo cubierto por pelos llamados sedas."},
	{ID: "2", Nombre: "Viudas Negras", Description: "tiene un cuerpo negro y brillante con una forma de reloj de arena rojo en la zona ventral"},
	{ID: "3", Nombre: "Licósidos", Description: "arañas que vagan en el suelo, excavando pequeñas galerías verticales u ocupando grietas naturales desde las que acechan a sus presas, cuya presencia detectan por las vibraciones del suelo."},
}

func getAranass(a *gin.Context) {
	a.IndentedJSON(http.StatusOK, aranass)
}

func getAranassByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range aranass {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
}
func main() {
	router := gin.Default()
	router.GET("/aranass", getAranass)
	router.GET("/aranass/:id", getAranassByID)
	router.Run("localhost:4000")

}

///Edinson
