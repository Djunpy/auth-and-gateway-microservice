package accounts

import db "auth-and-gateway-microservice/db/sqlc"

type User struct {
	Username  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
	Photo     string `json:"photo,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

type UserAddress struct {
	City       string `json:"city"`
	Street     string `json:"street,omitempty"`
	PostalCode int64  `json:"postal_code"`
}

type CreateChat struct {
}

func ResponseUserAddress(address db.Address) UserAddress {
	return UserAddress{
		City:       address.City,
		Street:     address.Street.String,
		PostalCode: address.PostalCode.Int64,
	}
}

func ResponseUserPhone(phone db.Phone) db.UserPhone {
	return db.UserPhone{
		Number:      phone.Number,
		CountryCode: phone.CountryCode,
	}
}
