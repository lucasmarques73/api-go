package JWTService

import (
	"api/Models"
	"api/Services/GetEnvData"
	"errors"
	"log"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	hoursInDay = 24
	daysInWeek = 7
)

// MyCustomClaims is my custom claims
type MyCustomClaims struct {
	ID int64 `json:"id"`
	jwt.StandardClaims
}

// MySigningKey is key JWT
var MySigningKey []byte

func init() {
	var s GetEnvData.Settings
	s = GetEnvData.GetEnvData(s)

	MySigningKey = []byte(s.JwtSecret)
}

// GetToken - Generate token
func GetToken(id int64) string {

	claims := MyCustomClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * hoursInDay * daysInWeek).Unix(),
			IssuedAt:  time.Now().Unix(),
			Id:        strconv.Itoa(int(id)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, _ := token.SignedString(MySigningKey)

	return tokenString
}

// IsTokenValid - Token is valid?
func IsTokenValid(val string) (int64, error) {
	token, err := jwt.ParseWithClaims(val, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySigningKey, nil
	})

	switch err.(type) {
	case nil:
		if !token.Valid {
			return 0, errors.New("Token is invalid")
		}

		var userID int64

		claims, ok := token.Claims.(*MyCustomClaims)

		if !ok {
			return 0, errors.New("Token is invalid")
		}

		userID = claims.ID

		return userID, nil
	case *jwt.ValidationError:
		vErr := err.(*jwt.ValidationError)

		switch vErr.Errors {
		case jwt.ValidationErrorExpired:
			return 0, errors.New("Token Expired, get a new one")
		default:
			log.Println(vErr)
			return 0, errors.New("Error while Parsing Token")
		}
	default:
		return 0, errors.New("Unable to parse token")
	}
}

// GetUserFromToken - Get data from a authenticated user
func GetUserFromToken(tokenVal string) (user Models.User, err error) {
	if tokenVal == "" {
		err = errors.New("No token present")
		return
	}

	userID, err := IsTokenValid(tokenVal)
	if err != nil {
		err = errors.New("Token is invalid")
		return
	}

	if userID < 1 {
		err = errors.New("Token missing required data")
		return
	}

	// var user Models.User
	res := Models.UsersModel.Find(userID)
	err = res.One(&user)
	if err != nil || user.ID < 1 {
		err = errors.New("User in token does not exist in system")
		return
	}

	return user, err
}
