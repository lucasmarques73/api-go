package GetEnvData

import (
	"api/Errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Settings app
type Settings struct {
	DbUser     string
	DbPassword string
	DbName     string
	DbHost     string
	DbPort     string
	DbLogging  bool
	JwtSecret  string
}

//GetEnvData -  Func to Get dot env data
func GetEnvData(s Settings) Settings {

	err := godotenv.Load()
	Errors.CheckErr(err)

	host := os.Getenv("DB_HOST")
	if len(host) > 0 {
		s.DbHost = host
	} else {
		s.DbHost = "localhost"
	}

	database := os.Getenv("DB_DATABASE")
	if len(database) > 0 {
		s.DbName = database
	} else {
		s.DbName = "database"
	}

	user := os.Getenv("DB_USER")
	if len(user) > 0 {
		s.DbUser = user
	} else {
		s.DbUser = "user"
	}

	password := os.Getenv("DB_PASSWORD")
	if len(password) > 0 {
		s.DbPassword = password
	} else {
		s.DbPassword = "password"
	}
	port := os.Getenv("DB_PORT")
	if len(port) > 0 {
		s.DbPort = port
	} else {
		s.DbPort = "5432"
	}

	logging := os.Getenv("DB_LOGGING")
	if len(logging) > 0 {
		s.DbLogging, _ = strconv.ParseBool(logging)
	} else {
		s.DbLogging = false
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if len(jwtSecret) > 0 {
		s.JwtSecret = jwtSecret
	} else {
		s.JwtSecret = "jwtSecret"
	}
	return s
}
