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

<<<<<<< HEAD
type Alimento struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
=======
// Lista de objetos de ejemplo

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
>>>>>>> f286d7655deeae2dfc17cced1077a98cc9b8fe4c
	Description string `json:"description"`
}

type Frutas struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Description string `json:"description"`
}

type Motos struct {
	ID          string  `json:"id"`
	Nombre      string `json:"nombre"`
	Description string `json:"description"`
}


var alimentos = []Alimento{
	{ID: 1, Nombre:"Arepas de choclo", Description: "Hechas de maíz"},
	{ID: 1, Nombre:"Arepitas", Description: "Hechas de maicitos"},
	{ID: 1, Nombre:"Buñuelos", Description: "redondos"},
}



func getAlimento(a *gin.Context){
a.IndentedJSON(http.StatusOK, alimentos)
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

var alimentos = []Alimento{
	{ID: 1, Nombre: "Arepas de choclo", Description: "Hechas de maíz"},
	{ID: 1, Nombre: "Arepitas", Description: "Hechas de maicitos"},
	{ID: 1, Nombre: "Buñuelos", Description: "redondos"},
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

func getAliimento(a *gin.Context){
a.IndentedJSON(http.StatusOK, alimentos)
	{ID: 4, Tipo: "Pulga de santi", Nombre: "En caballo"},

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


// LUISA VILLACORTE
var fruits = []Frutas{
	{ID: 1, Nombre: "Sandia", Description: "Fruta tropical"},
	{ID: 2, Nombre: "Limón", Description: "Fruta ácida"},
	{ID: 3, Nombre: "Naranja", Description: "Fruta semi-ácida"},
}
func getFrutas(a *gin.Context) {
	a.IndentedJSON(http.StatusOK, fruits)
}

func getFrutaByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID no válido"})
		return
	}

<<<<<<< HEAD
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
=======
	var frutaEncontrada Frutas
	for _, fruta := range fruits {
		if fruta.ID == id {
			frutaEncontrada = fruta
			break
		}
	}

	if frutaEncontrada.Nombre == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Fruta no encontrada"})
		return
	}

	c.JSON(http.StatusOK, frutaEncontrada)
>>>>>>> 2ddf426d11b1705288f4fb68e29c078de0e4fb05
}




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
func buscarFrutaPorID(idBuscado int) *Server {
	 for i := range frutas {
        if frutas[i].ID == idBuscado {
            return &frutas[i] // Devolver un puntero al objeto encontrado
        }
    }
	return nil // Retorna nil si no se encuentra la fruta
}


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
// Función handler para obtener una fruta por su ID

	// Convertir el ID de string a int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Buscar la fruta por ID
	frutaEncontrada := buscarFrutaPorID(id)

	// Verificar si la fruta fue encontrada
	if frutaEncontrada != nil {
		// Devolver la fruta encontrada como JSON
		c.JSON(http.StatusOK, frutaEncontrada)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Fruta no encontrada"})
	}
}
// LUISA VILLACORTE

<<<<<<< HEAD
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

func getOso(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, osos)
}

func getOsoPorID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID debe ser un número entero"})
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
// Yerson Fingit

func main() {
	router := gin.Default()
	router.GET("/aranas", getAranas)
	router.GET("/aranas/:id", getAranaByID)
	router.GET("/osos", getOso)
	router.GET("/osos/:id", getOsoPorID)
	router.Run("localhost:4000")
	router.GET("/frutamango/:id", getFruta)
	router.GET("/frutamango", getAll)
=======

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
func main(){
	router := gin.Default()

	router.GET("/aranas", getAranas)
	router.GET("/aranas/:id", getAranaByID)

	router.GET("/frutamango/:id", getFruta)
	router.GET("/frutamango", getAll)


	router.GET("/frutamango/:id", getFruta)
	router.GET("/frutamango", getAll)

>>>>>>> 2ddf426d11b1705288f4fb68e29c078de0e4fb05
	router.GET("/pulgasM", getPulga)
	router.GET("/pulgasM/:id", getByIdPulga)
	router.GET("/animales", getAnimales)
	router.GET("/animales/:id", getAnimal)


	router.GET("/frutas/:id", getFrutaByIDD)
	router.GET("/ave", getAve)
	router.GET("/ave/:id",getAveId)
<<<<<<< HEAD
=======
=======
	router.GET("/frutas/:id", getFrutaByIDD)
	router.GET("/motos", getMotos)
	router.GET("/motos/:id", getMotosByID)

	router.GET("/frutas", getFrutas)
	router.GET("/frutas/:id", getFrutaByID)
	router.GET("/ave", getAve)
	router.GET("/ave/:id",getAveId)

	router.GET("/frutas", getFrutas)
	router.GET("/frutas/:id", getFrutaByID)
	router.GET("/serpiente", getSerpiente)
	router.GET("/serpiente/:id", getSerpienteByID)
>>>>>>> 2ddf426d11b1705288f4fb68e29c078de0e4fb05
	router.Run("localhost:4000")

}



