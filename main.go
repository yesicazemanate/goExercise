package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
type AnimalAve struct{
ID int `json:"id"`
Nombre string `json:"nombre"`
Ubicacion string `json:"ubicacion"`
Habitat string `json:"habitat"`
}// yesica zemanate
// yesica zemanate
var ave =[]AnimalAve{
	{ID: 1, Nombre: "Águila calva",Ubicacion:"América del Norte", Habitat: " Bosques cercanos a cuerpos de agua, como ríos, lagos y costas marinas"},
	{ID: 2, Nombre :"Colibrí esmeralda ", Ubicacion: "América Central y del Sur", Habitat: "Bosques tropicales y subtropicales, jardines y áreas con flores"},
	{ID:3, Nombre:"Cóndor de los Andes ", Ubicacion: "Cordilleras de los Andes en América del Sur", Habitat: " Cordilleras montañosas, acantilados y cielos abiertos"},
	}// yesica zemanate
	func getAve(a *gin.Context){
	
a.IndentedJSON(http.StatusOK, ave)
}
func getAveId(a *gin.Context){
	ids:= a.Param("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		a.JSON(http.StatusNotFound, gin.H{"error": "ID no válido"})
		return
	}
	for _, c := range ave{
		if c.ID == id{
		a.IndentedJSON(http.StatusOK,c )	
		return
	}
}}// yesica zemanate

func main() {
	router := gin.Default()
	 router.GET("/ave", getAve)
	 router.GET("/ave/:id",getAveId)








	router.Run("localhost:4000")
}
