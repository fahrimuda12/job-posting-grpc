package main

import (
	"io"
	"job-posting/gateway/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.DisableConsoleColor()

   // Logging to a file.
   f, _ := os.Create("gin.log")
   gin.DefaultWriter = io.MultiWriter(f)
	router := gin.Default()
	routes.Routes(router)
	
	// //mulai server dengan port 3000
	router.Run(":4050")
}