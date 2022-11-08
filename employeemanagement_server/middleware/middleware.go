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

	admin:=[]string{"admin"}
	a := g.Group("/a")
	a.Use(IsAllowed(admin))

	hr:=[]string{"hr"}
	h := g.Group("/h")
	h.Use(IsAllowed(hr))

	emp:=[]string{"employee"}
	e := g.Group("/e")
	e.Use(IsAllowed(emp))

	adHr:=[]string{"admin","hr"}
	ah := g.Group("/ah")
	ah.Use(IsAllowed(adHr))

	empHr:=[]string{"employee","hr"}
	he := g.Group("/he")
	he.Use(IsAllowed(empHr))
	route.Init(r,o,a,h,ah,he,e)
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

func IsAllowed(role []string) gin.HandlerFunc { 
	return func(c *gin.Context) {
		checkAuth()
		login, _ := helpers.GetLoginFromToken(c)
		if contains(role,login.Role){
			c.Next()
		}else{
			fmt.Println("Unauthorized access")
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized access"})
		}
	}
}

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

