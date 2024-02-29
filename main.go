package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


//YERSON
type Oso struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}

var osos = []Oso{
	{ID: 1, Nombre: "Oso Negro Americano", Descripcion: "Habita en Norteamérica."},
	{ID: 2, Nombre: "Oso Polar", Descripcion: "Habita en el Ártico."},
	{ID: 3, Nombre: "Oso Panda", Descripcion: "Caracterizado por su pelaje blanco y negro."},
}

func getOso(c *gin.Context) {
    c.IndentedJSON(http.StatusOK,osos)
}
func getOsoPorID(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Error al convertir ID a entero"})
        return
    }

    for _, oso := range osos {
        if oso.ID == id {
            c.JSON(http.StatusOK, oso)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "Oso no encontrado"})
}


func main() {
	router := gin.Default()
	router.GET("/osos", getOso)
    router.GET("/osos/:id", getOsoPorID)
	router.Run("localhost:4000")
}