package userEntity

type User struct {
	Id          int32
	Name        string
	PhoneNumber string
}

func CreateUser(name, phoneNumber string) *User {
	return &User{Name: name, PhoneNumber: phoneNumber}
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) SetPhoneNumber(phoneNumber string) {
	u.PhoneNumber = phoneNumber
}

func (u *User) GetPhoneNumber() string {
	return u.PhoneNumber
}
