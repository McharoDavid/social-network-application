package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

//Address is used to hold the address details of the account user
type Address struct {
	HouseNumber string `json:"housenumber,omitempty"`
	Street      string `json:"street,omitempty"`
	City        string `json:"city,omitempty"`
	ZipCode     string `json:"zipcode,omitempty"`
}

// Value Make the Address struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (a Address) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Make the Address struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (a *Address) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
