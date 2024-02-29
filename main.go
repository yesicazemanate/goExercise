package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
//yesica zemanate
type AnimalAve struct{
ID string `json:"id"`
Nombre string `json:"nombre"`
Ubicacion string `json:"ubicacion"`
Habitat string `json:"habitat"`
}// yesica zemanate
type Alimento struct{
	ID int `json:"id"`
	Nombre string `json:"nombre"`
	Description string `json:"description"`
}

type Frutas struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Description string `json:"description"`
}


type Arana struct {
	ID           int     `json:ID`
	Tipo string `json:ID`
	Nombre string `json:Nombre`
}// yesica zemanate
var ave =[]AnimalAve{
	{ID: "1", Nombre: "Águila calva",Ubicacion:"América del Norte", Habitat: " Bosques cercanos a cuerpos de agua, como ríos, lagos y costas marinas"},
	{ID: "2", Nombre :"Colibrí esmeralda ", Ubicacion: "América Central y del Sur", Habitat: "Bosques tropicales y subtropicales, jardines y áreas con flores"},
	{ID: "3", Nombre:"Cóndor de los Andes ", Ubicacion: "Cordilleras de los Andes en América del Sur", Habitat: " Cordilleras montañosas, acantilados y cielos abiertos"},
	}// yesica zemanate
var alimentos = []Alimento{
	 {ID: 1, Nombre:"Arepas de choclo", Description: "Hechas de maíz"},
	{ID: 1, Nombre:"Arepitas", Description: "Hechas de maicitos"},
	{ID: 1, Nombre:"Buñuelos", Description: "redondos"},
}

var aranas = []Arana{
	{ID: 1, Tipo: "Tarantula", Nombre: "kidd keo"},
	{ID: 2, Tipo: "Aranita chiquita", Nombre: "Ariana Grande"},
}
// yesica zemanate
func getAve(a *gin.Context){
	
a.IndentedJSON(http.StatusOK, ave)
}
func getAveId(a *gin.Context){
	id:= a.Param("id")
	for _, c := range ave{
		if c.ID == id{
		a.IndentedJSON(http.StatusOK,c )	
		return
	}
}}// yesica zemanate

func getAlimento(a *gin.Context){
a.IndentedJSON(http.StatusOK, alimentos)
}


func getAranas(a *gin.Context) {
	a.IndentedJSON(http.StatusOK, aranas)
}

func getAranaByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID no válido"})
		return
	}

	var arana Arana
	for _, a := range aranas {
		if a.ID == id {
			arana = a
			break
		}
	}

	if arana.Nombre == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Arana no encontrada"})
		return
	}

	c.JSON(http.StatusOK, arana)
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
	// if frutaEncontrada.ID == nil {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "Fruta no encontrada"})
	// 	return
	// }
	
	// Devolver la fruta encontrada
	c.IndentedJSON(http.StatusOK, frutaEncontrada)
}


func main(){
	router := gin.Default()
	router.GET("/alimentos", getAlimento)
	router.GET("/aranas", getAranas)
	router.GET("/aranas/:id", getAranaByID)
	router.GET("/ave", getAve)
	router.GET("/ave/:id",getAveId)
	router.Run("localhost:4000")
}