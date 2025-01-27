package entities

type User struct {
	Id           int32
	Name         string
	Phone_number string
}

func CreateUser(_name string, _phone string) *User {
	return &User{Name: _name, Phone_number: _phone}
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) SetPhone(phone string) {
	u.Phone_number = phone
}

func (u *User) GetNumber() string {
	return u.Phone_number
}
