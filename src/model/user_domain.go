package model

type userDomain struct {
	id string
	email string
	password string
	name string
	age int8
}

func (user *userDomain) GetID() string {
	return user.id
}
func (user *userDomain) GetEmail() string {
	return user.email
}
func (user *userDomain) GetPassword() string {
	return user.password
}
func (user *userDomain) GetName() string {
	return user.name
}
func (user *userDomain) GetAge() int8 {
	return user.age
}
func (user *userDomain) SetID(id string) {
	user.id = id
}