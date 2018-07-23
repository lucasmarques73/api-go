package Controllers

import (
	"api/Models"
	"api/Services/JWTService"
	"api/Services/PasswordService"
	"encoding/json"
	"net/http"
)

type ResponseLoginData struct {
	Token string      `json:"token"`
	User  Models.User `json:"user"`
}

type LoginData struct {
	Email string `json:"email"`
	Pass  string `json:"password"`
}

// Login - User Login
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var ld LoginData
	_ = json.NewDecoder(r.Body).Decode(&ld)

	if len(ld.Email) < 1 || len(ld.Pass) < 1 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(Response{
			Errors:  true,
			Data:    "",
			Message: "Email and password are required.",
		})
		return
	}
	var user Models.User
	res := Models.UsersModel.Find("email", ld.Email)
	err := res.One(&user)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Response{
			Errors:  true,
			Data:    "",
			Message: "Email not found",
		})
		return
	}

	if !PasswordService.IsValid(user.Pass, ld.Pass) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(Response{
			Errors:  true,
			Data:    "",
			Message: "The password is invalid",
		})
		return
	}

	token := JWTService.GetToken(user.ID)
	login := ResponseLoginData{User: user, Token: token}

	json.NewEncoder(w).Encode(Response{
		Errors:  false,
		Data:    login,
		Message: "User Logged",
	})

}

// Logout User Logout
func Logout() {}

// RefreshToken - User Refresh Token
func RefreshToken() {}

// GetAuthUser - Get User Authenticated
func GetAuthUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	ua := r.Header.Get("Authorization")

	tokenString := ua[7:len(ua)]

	var user Models.User
	user, err := JWTService.GetUserFromToken(tokenString)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{
			Errors:  true,
			Data:    "",
			Message: err.Error(),
		})
	}

	json.NewEncoder(w).Encode(Response{
		Errors:  false,
		Data:    user,
		Message: "User that is logged in",
	})
}
