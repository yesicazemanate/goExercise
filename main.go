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
// func main () {
// router := gin.Default()
// router.GET("/alimentos", getAlimento)
// router.Run("localhost:4000")
// }

// package main
//  import(
// "github.com/gin-gonic/gin"
// "net/http"
// )
// type Frutas struct{
// 	ID int `json:"id"`
// 	Nombre string `json:"nombre"`
// 	Description string `json:"description"`
// }
// var fruits = []Frutas{
// 	{ID : 1 , Nombre: "Sandia" , Description: "Fruta tropical"},
// 	{ID : 2, Nombre: "Limón" , Description: "Fruta ácida"},
// 	{ID : 3, Nombre: "Naranja", Description: "Fruta semi-ácida"},
// }
// func getFruta(f*gin.Context)  {
// 	f.IndentedJSON(http.StatusOK, fruits)

// }
// func main () {
// router := gin.Default()
// router.GET("/frutas", getFruta)
// router.Run("localhost:5000")
// }

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Frutas struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Description string `json:"description"`
}

var fruits = []Frutas{
	{ID: 1, Nombre: "Sandia", Description: "Fruta tropical"},
	{ID: 2, Nombre: "Limón", Description: "Fruta ácida"},
	{ID: 3, Nombre: "Naranja", Description: "Fruta semi-ácida"},
}

func getFrutaByID(c *gin.Context) {
	idStr := c.Param("id") 
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de ID inválido"})
		return
	}

	// Encontrar la fruta por ID
	var frutaEncontrada *Frutas
	for _, c := range fruits {
		if c.ID == id {
			frutaEncontrada = &c
			break
			
			
		}
	}

	// Comprobar si se encontró la fruta
	if frutaEncontrada.ID == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Fruta no encontrada"})
		return
	}

	// Devolver la fruta encontrada
	c.IndentedJSON(http.StatusOK, frutaEncontrada)
}

func main() {
	router := gin.Default()
	router.GET("/frutas/:id", getFrutaByID) // Definir ruta con parámetro ":id"
	router.Run("localhost:6000")
}
