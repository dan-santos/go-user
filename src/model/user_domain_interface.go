package model

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetID() string
	GetAge() int8

	SetID(string)

	EncryptPassword()
}

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain {
		email: email,
		name: name,
		age: age,
		password: password,
	}
}