package handlers

import (
	"encoding/json"
	"github/social-network/models"
	"io/ioutil"
	"net/http"
)

//Login is used to login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var account models.Account
	//read the body of the request
	body, err := ioutil.ReadAll(r.Body)
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

	if account.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Email Required, Please Enter the Email"))
		return
	}
	if account.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Password Required, Please Enter the Password"))
		return
	}

	// Get Account details from DB using Email
	accs, err := dbmgr.GetAccount("", account.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// if length of accounts is zero, then account with email not found
	if len(accs) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No account found with given Email"))
		return
	}

	if accs[0].Password != account.Password {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid password"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login Successfully"))

}
