
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//Brayan Delgado
type Server struct {
    ID          int
    Nombre      string
    Description string
    
}



type Motos struct {
	ID          string  `json:"id"`
	Nombre      string `json:"nombre"`
	Description string `json:"description"`
}



////////SANTIAGO NARVAEZ LASSO

type Arana struct {
	ID     int    `json:ID`
	Tipo   string `json:ID`
	Nombre string `json:Nombre`
}


type Serpiente struct{
	ID int `json:ID`
	Tipo string `json:"Tipo"`
	Nombre string `json:"Nombre"`
}
type Animales struct {
	ID         int     `json:ID`
	Tipo string `json:ID`
	Nombre string `json:Nombre`
}



var aranas = []Arana{
	{ID: 1, Tipo: "Tarantula", Nombre: "kidd keo"},
	{ID: 2, Tipo: "Aranita chiquita", Nombre: "Ariana Grande"},
}
//Sthefany Rodriguez
var serpiente = []Serpiente{
	{ID: 1, Tipo: "Cobra Real", Nombre: "Ophiophagus"},
	{ID: 2, Tipo: "Mamba Negra ", Nombre: "Dendroaspis "},
}

func getSerpiente(s * gin.Context){
	s.IndentedJSON(http.StatusOK, serpiente)
}

func getSerpienteByID(s *gin.Context){
	idse := s.Param("id")
	idint, err  := strconv.Atoi(idse)
	if err != nil {
		s.IndentedJSON(http.StatusNotFound, gin.H{"message": "Serpiente no encontrada"})
	}

	//Sthefanny Rodriguez


	for _, e := range serpiente {
		if e.ID == idint{
			s.IndentedJSON(http.StatusOK, e)
			return
		}
	}
	
}
//Maryuri
var pulgas = []Animales{
	{ID: 1, Tipo: "Tungidae.", Nombre: "En casa"},
	{ID: 2, Tipo: "Pulga de perros", Nombre: "En caballo"},
	{ID: 3, Tipo: "Pulga de gatos", Nombre: "En caballo"},
	{ID: 4, Tipo: "Pulga de santi", Nombre: "En caballo"},}
// yesica zemanate


// yesica zemanate
// func getAve(a *gin.Context){
// a.IndentedJSON(http.StatusOK, ave)
// }
// func getAveId(a *gin.Context){
// 	id:= a.Param("id")
// 	for _, c := range ave{
// 		if c.ID == id{
// 		a.IndentedJSON(http.StatusOK,c )	
// 		return
// 	}
// }}


func getAranas(a *gin.Context) {
	a.IndentedJSON(http.StatusOK, aranas)
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
////////SANTIAGO NARVAEZ LASSO





//Darbin
type Animal struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Description string `json:"description"`
}

var animales = []Animal{
	{ID: 1, Nombre: "Leon verde", Description: "es muy grande"},
	{ID: 2, Nombre: "Leon rojo", Description: "es muy grande"},
	{ID: 3, Nombre: "Leon negro", Description: "es muy grande"},
	{ID: 4, Nombre: "Leon gris", Description: "es muy grande"},

}

func getAnimales(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, animales)
}

func getAnimal(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}


	var animal Animal
	for _, a := range animales {
		if a.ID == id {
			animal = a
			break
		}
	}


	if animal.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Animal no encontrado"})
		return
	}


	c.IndentedJSON(http.StatusOK, animal)
}
// BRAYAN DELGADO//
var frutas = []*Server{
	{ID: 1, Nombre: "Manzana", Description: "fruta que se da en climas frios"},
	{ID: 2, Nombre: "Uva", Description: "fruta dulce"},
	{ID: 3, Nombre: "Banano", Description: "fruta con mucha proteina"},
	{ID: 4, Nombre: "Mango", Description: "fruta tropical"},
}

// Función para buscar una fruta por su ID en una lista de frutas

//Andrea salazar perez
type Fruta struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Description string `json:"description"`
}

var mangos = []Fruta{
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

	var mango Fruta
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

//Andrea salazar perez



//KEVIN //
var motos = []Motos{
	{ID: "1", Nombre: "yamaha", Description: "velocidad 1"},
	{ID: "2", Nombre: "honda", Description: "velocidad 2"},
	{ID: "3", Nombre: "cbr", Description: "velocidad 3"},
}

func getMotos(a *gin.Context) {
	a.IndentedJSON(http.StatusOK, motos)
}

func getMotosByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range motos {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
}

//jossie aranda 
type Panda struct{
	ID string `json:"id"`
	Nombre string `json:"nombre"`
	Description string `json:"description"`
}

var animal = []Panda{
	 {ID: "1", Nombre:"oso", Description: "es grande "},
	{ID:"2" , Nombre:"perro", Description: "son bravos"},
	{ID: "3", Nombre:"gato", Description: "tienen grandes uñas"},
}
func getanimal(a *gin.Context){
a.IndentedJSON(http.StatusOK, animal)

}
func getAnimalByID(c *gin.Context) {
	
	 id := c.Param("id") 
	for _,a := range animal{
		if a.ID == id{
			c.IndentedJSON(http.StatusOK,a)
			return
		}
	}
	
}
//jossie aranda 
	


// Yerson
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



///edinson
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



///edinson
// Yerson Fingit

func main() {
	router := gin.Default()
	router.GET("/aranas", getAranas)
	router.GET("/aranas/:id", getAranaByID)
	router.GET("/frutamango/:id", getFruta)
	router.GET("/frutamango", getAll)
	router.GET("/pulgasM", getPulga)
	router.GET("/pulgasM/:id", getByIdPulga)
	router.GET("/animales", getAnimales)
	router.GET("/animales/:id", getAnimal)
	router.GET("/aranass", getAranass)
	router.GET("/aranass/:id", getAranassByID)
	router.GET("/ave", getAve)
	router.GET("/ave/:id",getAveId)
    router.GET("/alimentos", getAlimento)
    router.GET("/animal/:id", getAnimalByID)
	router.GET("/motos", getMotos)
	router.GET("/motos/:id", getMotosByID)
	router.GET("/ave", getAve)
	router.GET("/ave/:id",getAveId)
	router.Run("localhost:4000")
    router.GET("/animal/:id", getAnimalByID)
	router.GET("/motos", getMotos)
	router.GET("/motos/:id", getMotosByID)
	router.GET("/serpiente", getSerpiente)
	router.GET("/serpiente/:id", getSerpienteByID)
	router.Run("localhost:4000")

}




