package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Mango struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Description string `json:"description"`
}

var mangos = []Mango{
	{ID: 1, Nombre: "Mango verde", Description: "Es un mango que todavía está verde y no es jugoso"},
	{ID: 2, Nombre: "Mango rojo", Description: "Es un mango que es muy jugoso ya que su color rojo lo demuestra"},
	{ID: 3, Nombre: "Mango amarillo", Description: "Es un mango que está entre jugoso y verde"},
}

func getFruta(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID debe ser un número entero"})
		return
	}

	var mango Mango
	for _, item := range mangos {
		if item.ID == id {
			mango = item
			break
		}
	}

	if mango.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mango no encontrado"})
		return
	}

	c.IndentedJSON(http.StatusOK, mango)
}

func getAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, mangos)
}


func main() {
	router := gin.Default()
	router.GET("/frutamango/:id", getFruta)
	
	router.GET("/frutamango", getAll)
	router.Run("localhost:4000")
}


