package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


//maryuri
type Animales struct {
	ID  int  `json:ID`
	Tipo string `json:ID`
	Nombre string `json:Nombre`
}
//Maryuri
var pulgas = []Animales{
	{ID: 1, Tipo: "Tungidae.", Nombre: "En casa"},
	{ID: 2, Tipo: "Pulga de perros", Nombre: "En caballo"},
	{ID: 3, Tipo: "Pulga de gatos", Nombre: "En caballo"},
	{ID: 4, Tipo: "Pulga de santi", Nombre: "En caballo"},
}
//Get Maryuri
func getPulga(m *gin.Context) {
	m.IndentedJSON(http.StatusOK, pulgas)
}
//GET by Id MARYURI
func getByIdPulga(m * gin.Context){
	idStr := m.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
			m.IndentedJSON(http.StatusNotFound, gin.H{"message": "id not found"})
	}
	for _, ma := range pulgas{
		if ma.ID == id{
			m.IndentedJSON(http.StatusOK, ma)
			return
		}
	}

}

func main() {
	router := gin.Default()

	router.GET("/pulgasM", getPulga)
router.GET("/pulgasM/:id", getByIdPulga)


	router.Run("localhost:4000")
}
     