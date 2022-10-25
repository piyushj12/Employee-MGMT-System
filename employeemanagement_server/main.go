package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	database "EmployeeAssisgnment/api/database"
	"github.com/gin-contrib/cors"
	middleware "EmployeeAssisgnment/middleware"

)

func main() {
	fmt.Println("Gin-Gonic Server")
	//database.InitDB()
	startServer()
}

// func startServer() {
// 	router := gin.Default()
// 	router.GET("/", checkStatus())
// 	route.Init(router)
// 	s := &http.Server{
// 		Addr:    ":4700",
// 		Handler: router,
// 	}
// 	s.ListenAndServe()
// }

func checkStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "Server is running successfully !!!!!")
	}
}

func startServer() {
	router := gin.Default()
	md := cors.DefaultConfig()
	md.AllowAllOrigins = true
	md.AllowHeaders = []string{"*"}
	md.AllowMethods = []string{"*"}
	router.Use(cors.New(md))
	middleware.InitMiddleware(router)
	router.GET("/", checkStatus())

	s := &http.Server{
		Addr:    ":4700",
		Handler: router,
	}
	database.InitDB()
	s.ListenAndServe()
}
