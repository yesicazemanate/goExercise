
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// edinson
type Spider struct {
	ID          string `json:"id"`
	Nombre      string `json:"nombre"`
	Description string `json:"description"`
}

var spider = []Spider{
	{ID: "1", Nombre: "Tarantulas", Description: "arañas migalomorfas de gran tamaño con el cuerpo cubierto por pelos llamados sedas."},
	{ID: "2", Nombre: "Viudas Negras", Description: "tiene un cuerpo negro y brillante con una forma de reloj de arena rojo en la zona ventral"},
	{ID: "3", Nombre: "Licósidos", Description: "arañas que vagan en el suelo, excavando pequeñas galerías verticales u ocupando grietas naturales desde las que acechan a sus presas, cuya presencia detectan por las vibraciones del suelo."},
}

func getSpider(a *gin.Context) {
	a.IndentedJSON(http.StatusOK, spider)
}

func getSpiderByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range spider {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
}
func main() {
	router := gin.Default()
	router.GET("/spider", getSpider)
	router.GET("/spider/:id", getSpiderByID)
	router.Run("localhost:4000")

}

///Edinson
