package handlers

import (
	"encoding/json"
	"errors"
	"github/social-network/models"
	"io/ioutil"
	"net/http"

	"github.com/google/logger"
	"github.com/gorilla/mux"
)

//UpdateAccount is used to update account details of user
func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	//read url params of the request
	urlParams := mux.Vars(r)

	var changes models.Account

	//read the body of the request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if err = json.Unmarshal(body, &changes); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	//Get Account details from DB using ID
	accs, err := dbmgr.GetAccount(urlParams["id"], "")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if len(accs) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No account found with the given ID"))
		return
	}

	//update account
	isValidChange, err := applyChanges(&accs[0], &changes)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// if no changes detected, then return error
	if isValidChange == false {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No changes detected"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Changes Made/Updated Successfully"))
	return
}

func applyChanges(account, changes *models.Account) (isValidChange bool, err error) {
	if changes.FirstName != "" && changes.FirstName != account.FirstName {
		account.FirstName = changes.FirstName
		isValidChange = true
	}
	if changes.MiddleName != "" && changes.MiddleName != account.MiddleName {
		account.MiddleName = changes.MiddleName
		isValidChange = true
	}
	if changes.LastName != "" && changes.LastName != account.LastName {
		account.LastName = changes.LastName
		isValidChange = true
	}
	if changes.Email != "" && changes.Email != account.Email {
		account.Email = changes.Email
		isValidChange = true
	}
	if changes.PhoneNumber != "" && changes.PhoneNumber != account.PhoneNumber {
		account.PhoneNumber = changes.PhoneNumber
		isValidChange = true
	}
	if changes.Password != "" && changes.Password != account.Password {
		if changes.OldPassword == "" || changes.OldPassword != account.Password {
			return false, errors.New("incorrect old password")
		}

		account.Password = changes.Password
		isValidChange = true
	}
	if changes.Gender != "" && changes.Gender != account.Gender {
		account.Gender = changes.Gender
		isValidChange = true
	}
	if changes.Status != "" && changes.Status != account.Status {
		account.Status = changes.Status
		isValidChange = true
	}
	if changes.Address != nil {
		if account.Address == nil {
			account.Address = &models.Address{}
		}

		if changes.Address.HouseNumber != "" && changes.Address.HouseNumber != account.Address.HouseNumber {
			account.Address.HouseNumber = changes.Address.HouseNumber
			isValidChange = true
		}
		if changes.Address.Street != "" && changes.Address.Street != account.Address.Street {
			account.Address.Street = changes.Address.Street
			isValidChange = true
		}
		if changes.Address.City != "" && changes.Address.City != account.Address.City {
			account.Address.City = changes.Address.City
			isValidChange = true
		}
		if changes.Address.ZipCode != "" && changes.Address.ZipCode != account.Address.ZipCode {
			account.Address.ZipCode = changes.Address.ZipCode
			isValidChange = true
		}
	}
	if isValidChange {
		// save account changes to DB
		err = dbmgr.SaveAccount(account)
		logger.Infof("changes applied for account %v", account)
	}
	return isValidChange, err
}
