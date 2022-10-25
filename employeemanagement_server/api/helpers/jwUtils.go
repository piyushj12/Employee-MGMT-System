package helpers

import (
	"EmployeeAssisgnment/api/model"
	"EmployeeAssisgnment/api/services"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var GlobalJWTKey string

func init() {
	GlobalJWTKey = "Educ@t!on123"
}

type jwtCustomClaim struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Role    string `json:"role"`
	EmpID    string `json:"empid"`
	jwt.StandardClaims
}

func GenerateToken(login model.Login, expirationTime time.Duration) (string, error) {
	claims := jwtCustomClaim{
		Username: login.Username,
		Password: login.Password,
		Role:login.Role,
		EmpID:login.EmpID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expirationTime).Unix(),
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(GlobalJWTKey))
	if err != nil {
		fmt.Println("Error during token generation", err)
	}
	return t, nil
}

func GetUserFromToken(c *gin.Context) model.Login {
	login, err := GetLoginFromToken(c)
	if err != nil {
		fmt.Println("Error : ", err)
		return model.Login{}
	}
	emp, isValid := services.ValidateUser(login)
	if !isValid {
		return model.Login{}
	}
	return emp
}

// GetLoginFromToken login object from JWT Token
func GetLoginFromToken(c *gin.Context) (model.Login, error) {

	login := model.Login{}
	decodedToken, err := DecodeToken(c.GetHeader("Authorization"), GlobalJWTKey)
	if err != nil {
		return login, errors.New("GetLoginFromToken - unable to decode token")
	}
	// login ID is the compulsary field, so haven't added check for nil
	if decodedToken["username"] == nil || decodedToken["username"] == "" {
		return login, errors.New("GetLoginFromToken - login id not found")
	}
	login.Username = decodedToken["username"].(string)
	login.Password = decodedToken["password"].(string)
	login.Role = decodedToken["role"].(string)
	login.EmpID = decodedToken["empid"].(string)
	return login, nil
}

func DecodeToken(tokenFromRequest, jwtKey string) (jwt.MapClaims, error) {

	// get data i.e.Claims from token
	token, err := jwt.Parse(tokenFromRequest, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtKey), nil
	})

	if err != nil {
		fmt.Println("Error while parsing JWT Token: ", err)
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Error while getting claims")
	}
	return claims, nil
}
