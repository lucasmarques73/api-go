package JWTService

import (
	"api/Models"
	"errors"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	HOURS_IN_DAY = 24
	DAYS_IN_WEEK = 7
)

type MyCustomClaims struct {
	id int64 `json:"id"`
	jwt.StandardClaims
}

var MySigningKey = []byte("K0xOQwFEr4WDgRW")

func GetToken(id int64) string {

	claims := MyCustomClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * HOURS_IN_DAY * DAYS_IN_WEEK).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, _ := token.SignedString(MySigningKey)

	return tokenString
}

func IsTokenValid(val string) (int64, error) {
	token, err := jwt.Parse(val, func(token *jwt.Token) (interface{}, error) {
		return MySigningKey, nil
	})

	switch err.(type) {
	case nil:
		if !token.Valid {
			return 0, errors.New("Token is invalid")
		}

		var userId int64

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return 0, errors.New("Token is invalid")
		}

		userId = int64(claims["id"].(float64))

		return userId, nil
	case *jwt.ValidationError:
		vErr := err.(*jwt.ValidationError)

		switch vErr.Errors {
		case jwt.ValidationErrorExpired:
			return 0, errors.New("Token Expired, get a new one.")
		default:
			log.Println(vErr)
			return 0, errors.New("Error while Parsing Token!")
		}
	default:
		return 0, errors.New("Unable to parse token")
	}
}

func GetUserFromToken(tokenVal string) (user Models.User, err error) {
	if tokenVal == "" {
		err = errors.New("No token present.")
		return
	}

	userId, err := IsTokenValid(tokenVal)
	if err != nil {
		err = errors.New("Token is invalid.")
		return
	}

	if userId < 1 {
		err = errors.New("Token missing required data.")
		return
	}

	// var user Models.User
	res := Models.UsersModel.Find(userId)
	err = res.One(&user)
	if err != nil || user.ID < 1 {
		err = errors.New("User in token does not exist in system.")
		return
	}

	return user, err
}