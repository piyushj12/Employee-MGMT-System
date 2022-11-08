package middleware

import (
	"fmt"
	"EmployeeAssisgnment/api/helpers"
	"EmployeeAssisgnment/api/route"
	"EmployeeAssisgnment/api/services"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Init -Init
func InitMiddleware(g *gin.Engine) {
	g.Use(cors.Default()) // CORS request
	fmt.Println("InitMiddleware called")

	o := g.Group("/o")
	o.Use(OpenRequestMiddleware())

	r := g.Group("/r")
	r.Use(RestrictedRequestMiddleware())

	a := g.Group("/a")
	a.Use(IsAdmin())

	// h := g.Group("/h")
	// h.Use(IsHr())

	// e := g.Group("/e")
	// e.Use(IsEmployee())

	// ah := g.Group("/ah")
	// ah.Use(IsAdminandHr())

	// he := g.Group("/he")
	// he.Use(IsHrandEmp())
	route.Init(r,o,a)
}

// Need to check JWT token here
func RestrictedRequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//fmt.Println("HEADER  ",c.Request.Header)
		token := c.GetHeader("Authorization")
		login, err := helpers.GetLoginFromToken(c)
		if err != nil {
			fmt.Println("Token not available", err)
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid API token"})
		}
		if strings.Trim(token, "") == "" {
			fmt.Println("Token not available")
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid API token"})
		}

		_, isValid := services.ValidateUser(login)
		if !isValid {
			fmt.Println("Failed to validate user")
			c.AbortWithStatusJSON(401, gin.H{"error": "Failed to validate user"})
		}
		c.Next()
		//fmt.Println("RestrictedRequestMiddleware called")
	}

}

func checkAuth() gin.HandlerFunc{
	return func(c *gin.Context) {
		
		token := c.GetHeader("Authorization")
		login, err := helpers.GetLoginFromToken(c)
		if err != nil {
			fmt.Println("Token not available", err)
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid API token"})
		}
		if strings.Trim(token, "") == "" {
			fmt.Println("Token not available")
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid API token"})
		}

		_, isValid := services.ValidateUser(login)
		if !isValid {
			fmt.Println("Failed to validate user")
			c.AbortWithStatusJSON(401, gin.H{"error": "Failed to validate user"})
		}
	}
}

func OpenRequestMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		fmt.Println("OpenRequestMiddleware called")
	}
}

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		checkAuth()
		login, _ := helpers.GetLoginFromToken(c)
		if login.Role=="admin"{
			c.Next()
		}else{
			fmt.Println("Unauthorized access")
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized access"})
		}
	}
}


