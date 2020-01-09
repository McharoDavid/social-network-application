package handlers

import (
	"encoding/json"
	"net/http"
)

//ListAccounts is used to get the list of account details
func ListAccounts(w http.ResponseWriter, r *http.Request) {

	var id string

	//capture the id from query params/parameters
	value, ok := r.URL.Query()["id"]
	if ok {
		id = value[0]
	}

	//get accounts or account from db
	accs, err := dbmgr.GetAccount(id, "")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	//Encode the response to json
	JSONResp, err := json.Marshal(&accs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	//send the response
	w.WriteHeader(http.StatusOK)
	w.Write(JSONResp)
}
