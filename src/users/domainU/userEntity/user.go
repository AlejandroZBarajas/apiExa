package userEntity

type User struct {
	ID    int32
	Name  string
	Phone string
}

func CreateUser(Name, Phone string) *User {
	return &User{Name: Name, Phone: Phone}
}
func (u *User) SetPhoneNumber(phone_number string) {
	u.Phone = phone_number
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) GetPhoneNumber() string {
	return u.Phone
}

func (u *User) GetName() string {
	return u.Name
}
