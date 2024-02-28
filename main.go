// package main
// import(
// "github.com/gin-gonic/gin"
// "net/http"
// )

// type Alimento struct{
// 	ID int `json:"id"`
// 	Nombre string `json:"nombre"`
// 	Description string `json:"description"`
// }

// var alimentos = []Alimento{
// 	 {ID: 1, Nombre:"Arepas de choclo", Description: "Hechas de maíz"},
// 	{ID: 1, Nombre:"Arepitas", Description: "Hechas de maicitos"},
// 	{ID: 1, Nombre:"Buñuelos", Description: "redondos"},
// }
// func getAlimento(a *gin.Context){
// a.IndentedJSON(http.StatusOK, alimentos)
// }
// func main(){
// router := gin.Default()
// router.GET("/alimentos", getAlimento)
// router.Run("localhost:4000")
//

// /Edinson
package main

import (
	"net/http"
	

	"github.com/gin-gonic/gin"
)
type Arana struct{
	ID string `json:"id"`
	Nombre string `json:"nombre"`
	Description string `json:"description"`
}

var aranas = []Arana{
	{ID: "1", Nombre: "Tarantulas", Description: "arañas migalomorfas de gran tamaño con el cuerpo cubierto por pelos llamados sedas."},
	{ID: "2", Nombre: "Viudas Negras", Description:"tiene un cuerpo negro y brillante con una forma de reloj de arena rojo en la zona ventral"},
	{ID: "3", Nombre: "Licósidos", Description:"arañas que vagan en el suelo, excavando pequeñas galerías verticales u ocupando grietas naturales desde las que acechan a sus presas, cuya presencia detectan por las vibraciones del suelo."},
}
func getArana(a*gin.Context){
	a.IndentedJSON(http.StatusOK,aranas)
}

func getAranaByID(c *gin.Context){
	id := c.Param ("id")
	for _,a := range aranas{
		if a.ID == id{
			c.IndentedJSON(http.StatusOK,a)
			return
		}
	}
}
	


func main(){
	router:=gin.Default()
	router.GET("/aranas",getArana)
	router.GET("/aranas/:id", getAranaByID)
	router.Run("localhost:5000")
	
}