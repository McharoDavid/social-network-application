package models

//Account is the data model of the social-network service
type Account struct {
	ID          string   `json:"id,omitempty"`
	FirstName   string   `json:"firstname,omitempty"`
	MiddleName  string   `json:"middlename,omitempty"`
	LastName    string   `json:"lastname,omitempty"`
	Email       string   `json:"email,omitempty"`
	PhoneNumber string   `json:"phonenumber,omitempty"`
	Password    string   `json:"password,omitempty"`
	Gender      string   `json:"gender,omitempty"`
	Status      string   `json:"status,omitempty"`
	Address     *Address `json:"address,omitempty"`
	OldPassword string   `json:"oldpassword,omitempty"`
}
