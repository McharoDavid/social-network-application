package handlers

import (
	"fmt"
	"github/social-network/models"
	"net/http"

	"github.com/gorilla/mux"
)

//DeleteAccount is used to delete the account
func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)

	//Get Account details from DB using ID
	accs, err := dbmgr.GetAccount(urlParams["id"], "")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if len(accs) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No account found with given ID"))
		return
	}

	// update account status to deleted
	accs[0].Status = models.StatusDeleted

	if err = dbmgr.SaveAccount(&accs[0]); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(fmt.Sprintf("Account %v deleted successfully\n", urlParams["id"])))
}
