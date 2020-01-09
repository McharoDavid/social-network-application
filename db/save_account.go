package db

import (
	"github/social-network/models"

	"github.com/google/logger"
)

//SaveAccount is used to save account details into DB
func (c *dbClient) SaveAccount(account *models.Account) error {
	accs, err := c.GetAccount(account.ID, "")
	if err != nil {
		return err
	}
	logger.Infof("saving %v account details into db", account.ID)

	if len(accs) == 0 {
		_, err = c.db.Exec("INSERT INTO accounts(id,firstname,middlename,lastname,email,password,phonenumber,gender,address,status) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10);",
			account.ID, account.FirstName, account.MiddleName, account.LastName, account.Email, account.Password, account.PhoneNumber, account.Gender, account.Address, account.Status)
	} else {
		_, err = c.db.Exec("UPDATE accounts SET firstname=$1,middlename=$2,lastname=$3,email=$4,password=$5,phonenumber=$6,gender=$7,address=$8,status=$9 WHERE id=$10;",
			account.FirstName, account.MiddleName, account.LastName, account.Email, account.Password, account.PhoneNumber, account.Gender, account.Address, account.Status, account.ID)
	}
	return err
}
