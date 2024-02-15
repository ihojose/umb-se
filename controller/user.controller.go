package controller

import (
	"airbusexpert/model"
	"airbusexpert/utils"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func UserLogin(rw http.ResponseWriter, req *http.Request) {
	var user model.User

	// Password Hash
	hash := sha256.New()
	hash.Write([]byte(req.FormValue("password")))
	passwd := hex.EncodeToString(hash.Sum(nil))

	// Get User Data
	result := db.Model(&user).First(&user, "id", req.FormValue("username"))

	if req.FormValue("username") != strconv.Itoa(int(user.ID)) || user.HashPass != passwd {
		log.Printf("User %v cannot be logged in!", user.ID)
		utils.SetResponseCode(req, rw, model.Response{
			Status:  http.StatusUnauthorized,
			Message: "Username or password is incorrect!",
		})
		return
	}

	errors.Is(result.Error, gorm.ErrRecordNotFound)

	log.Printf("User %v logged in...", user.ID)
	utils.SetResponse(req, rw, model.Token{
		Token: utils.JwtBuilder(user),
	})
}

func UpdateUser(rw http.ResponseWriter, req *http.Request) {
	var user model.User
	var data model.User

	// Get Json Requests
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&user); err != nil {
		panic(err)
		return
	}

	// Database registration
	db.First(&data).Where("id", user.ID)

	// User Entity
	data.Name = strings.ToUpper(user.Name)
	data.Surname = strings.ToUpper(user.Surname)
	data.IdType = user.IdType
	data.Role = user.Role

	db.Save(&data).Where("id", user.ID)

	// Response
	log.Println("The account was updated...")
	utils.SetResponse(req, rw, model.Response{
		Status:  http.StatusOK,
		Message: "The account was updated!",
	})
}

func RegisterUser(rw http.ResponseWriter, req *http.Request) {
	var user model.User

	// Get Json Requests
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&user); err != nil {
		panic(err)
		return
	}

	// Password process
	hash := sha256.New()
	hash.Write([]byte(user.Password))
	passwd := hex.EncodeToString(hash.Sum(nil))
	hashKey := sha256.New()
	hashKey.Write([]byte(fmt.Sprintf("%v", user.ID)))
	key := hex.EncodeToString(hashKey.Sum(nil))

	// User Entity
	user.Name = strings.ToUpper(user.Name)
	user.Surname = strings.ToUpper(user.Surname)
	user.HashPass = passwd
	user.HashKey = key
	user.HashToken = "-"

	// Database registration
	db.Create(&user)

	// Response
	log.Println("New user account registered...")
	utils.SetResponse(req, rw, model.Response{
		Status:  http.StatusOK,
		Message: "You registered!",
	})
}
