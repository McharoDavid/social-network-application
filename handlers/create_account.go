package handlers

import (
	"encoding/json"
	"github/social-network/models"
	"io/ioutil"
	"net/http"

	"github.com/google/logger"
	"github.com/google/uuid"
)

var (
	accounts []models.Account
)

//CreateAccount is used to create account
func CreateAccount(w http.ResponseWriter, req *http.Request) {
	//decode json data to our data model
	var account models.Account

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if err = json.Unmarshal(body, &account); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if account.FirstName == "" || account.LastName == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("First Name and Last Name are required"))
		return
	}
	//generate new id for the account
	account.ID = uuid.New().String()

	//Add status of account
	account.Status = models.StatusActive

	//Save account to DB
	if err = dbmgr.SaveAccount(&account); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	logger.Infof("account %v created successfully", account.ID)

	w.WriteHeader(http.StatusCreated)
	return
}
