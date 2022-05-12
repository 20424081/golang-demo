package Models


import (
	"time"
	// "goapp/Config"
)

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password []byte `json:"-"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (b *User) TableName() string {
	return "users"
}

func CreateAUser(user *User) (err error) {
	return nil
}

func GetUsers()(err error){
	return nil
}

func GetAUser(user *User, id string)(err error){
	return nil
}

func UpdateAUser(user *User, id string)(err error){
	return nil
}
func DeleteAUser(user *User, id string)(err error){
	return nil
}