package Models


import (
	"time"
	"goapp/Config"
)

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


func CreateAUser(user *User) (err error) {
}

func GetUsers()(err error){

}

func GetAUser(user *User, id string)(err error){

}

func UpdateAUser(user *User, id string)(err error){

}
func DeleteAUser(user *User, id string)(err error){

}