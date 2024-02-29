package main

import(
"github.com/gin-gonic/gin"
"net/http"
"strconv"
)



//YERSON
func main(){
	router:=gin.Default()



	router.Run("localhost:4000")
}