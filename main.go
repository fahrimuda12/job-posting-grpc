package main

import (
	"io"
	"job-posting/gateway/routes"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(1, 5)

// Middleware to check the rate limit.
func rateLimiter(c *gin.Context) {
    if !limiter.Allow() {
        c.JSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
        c.Abort()
        return
    }
    c.Next()
}


func main() {
	gin.DisableConsoleColor()

   // Logging to a file.
   f, _ := os.Create("gin.log")
   gin.DefaultWriter = io.MultiWriter(f)
	router := gin.Default()

	// use rate limiter
	router.Use(rateLimiter)

	// use cors
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET","POST","PUT","DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
		  return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	 }))

	//  membuat rate limiter
	// limiter := tollbooth.NewLimiter(1, nil)

	routes.Routes(router)
	
	// //mulai server dengan port 3000
	router.Run(":4050")
}