package models

type User struct{
	ID 		 int 	`json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	PhoneNum string `json:"phone_num"`
}
