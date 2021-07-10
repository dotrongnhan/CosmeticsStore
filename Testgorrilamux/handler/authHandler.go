package handler

import (
	"Testgorillamux/database"
	"Testgorillamux/models"
	"Testgorillamux/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type ErrorResponse struct {
	Err string
}

// var cookieHandler = securecookie.New(
// 	securecookie.GenerateRandomKey(64),
// 	securecookie.GenerateRandomKey(32))

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	stmt, err := database.DB.Prepare("INSERT INTO users(full_name, email, password, phone, address, date_of_birth, gender, avatar, role_id) VALUES (?,?,?,?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	var data map[string]string
	json.Unmarshal(body, &data)

	if data["password"] != data["password_confirm"] {
		err := ErrorResponse{
			Err: "Password does not match",
		}
		json.NewEncoder(w).Encode(err)
		return
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	if err != nil {
		fmt.Println(err)
		err := ErrorResponse{
			Err: "Password Encryption failed",
		}
		json.NewEncoder(w).Encode(err)
	}
	roleId := 2
	user := models.User{
		FullName:    data["full_name"],
		Email:       data["email"],
		Password:    pass,
		Phone:       data["phone"],
		Address:     data["address"],
		DateOfBirth: data["date_of_birth"],
		Gender:      data["gender"],
		Avatar:      data["avatar"],
		RoleId:      uint(roleId),
	}

	_, err = stmt.Exec(user.FullName, user.Email, user.Password, user.Phone, user.Address, user.DateOfBirth, user.Gender, user.Avatar, user.RoleId)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New user was created")
	json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	var data map[string]string
	json.Unmarshal(body, &data)
	// json.Unmarshal(body, &user)
	result, _ := database.DB.Query("SELECT full_name, email, password, phone, address, date_of_birth, gender, avatar, role_id FROM users WHERE email = ?", data["email"])
	var user models.User
	for result.Next() {
		err := result.Scan(&user.FullName, &user.Email, &user.Password, &user.Phone, &user.Address, &user.DateOfBirth, &user.Gender, &user.Avatar, &user.RoleId)
		if err != nil {
			panic(err.Error())
		}
	}
	if data["email"] != user.Email {
		fmt.Println(err)
		err := ErrorResponse{
			Err: "User not found",
		}
		json.NewEncoder(w).Encode(err)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		fmt.Println(err)
		err := ErrorResponse{
			Err: "Wrong password",
		}
		json.NewEncoder(w).Encode(err)
		return
	}
	token, err := util.GenerateJwt(strconv.Itoa(int(user.RoleId)))
	if err != nil {
		panic(err.Error())
	}
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}
	w.Header().Set("JWT", token)
	http.SetCookie(w, cookie)
	parsedToken, _ := util.ParseJwt(token)
	json.NewEncoder(w).Encode(parsedToken)
	json.NewEncoder(w).Encode("Login successfully")
	json.NewEncoder(w).Encode(user)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
		MaxAge:   -1,
	}
	http.SetCookie(w, cookie)
	json.NewEncoder(w).Encode("Success")
}
