package exercises

type User struct {
	Name     string
	Age      uint8
	Email    string
	Password string
}

func (u *User) ChangeName(newName string) {
	u.Name = newName
}

func (u *User) ChangeAge(newAge uint8) {
	u.Age = newAge
}

func (u *User) ChangeEmail(newEmail string) {
	u.Email = newEmail
}

func (u *User) ChangePassword(newPassword string) {
	u.Password = newPassword
}
