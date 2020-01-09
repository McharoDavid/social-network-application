package db

import (
	"fmt"
	"github/social-network/models"

	"github.com/google/logger"
)

//GetAccount is used to get/retrieve Account details from DB
func (c *dbClient) GetAccount(id string, email string) ([]models.Account, error) {
	logger.Infof("retrieving account from db with filter: %v %v", id, email)

	sqlQuery := "SELECT * FROM accounts"
	if id != "" {
		sqlQuery = fmt.Sprintf("%v WHERE id='%v'", sqlQuery, id)
	} else if email != "" {
		sqlQuery = fmt.Sprintf("%v WHERE email='%v'", sqlQuery, email)
	}
	logger.Infof("get account details from dbusing query: %v", sqlQuery)

	var accounts []models.Account

	rows, err := c.db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var account models.Account
		err = rows.Scan(&account.ID, &account.FirstName, &account.MiddleName, &account.LastName, &account.Email, &account.Password, &account.PhoneNumber, &account.Gender, &account.Address, &account.Status)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}
