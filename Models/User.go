package Models


import (
	"time"
	"goapp/Config"
)

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name" validate:"required,max=60,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password []byte `json:"-" validate:"required"`
	RefreshToken string `json:"-" validate:"max=255"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (b *User) TableName() string {
	return "users"
}

func CreateAUser(user *User) error {
	if err := Config.DB.Create(user).Error; err !=nil{
		return err
	}
	return nil
}

// func GetUsers(users *[]User)(err error){
	
// 	return nil
// }

func GetAUser(user *User, id string) error{
	if err := Config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func CountByEmail(email string) (count int64, err error){
	if err = Config.DB.Model(&User{}).Where("email = ? ", email).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
func GetAUserByEmail(user *User, email string) error {
	if err := Config.DB.Where("email = ?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func GetAUserForRefresh(user *User, id uint, refreshToken string) error {
	if err := Config.DB.Where(&User{Id: id, RefreshToken: refreshToken}).First(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAUser(user *User) error {
	if err := Config.DB.Save(user).Error; err != nil{
		return err
	}
	return nil
}
func DeleteAUser(user *User, id string) error {
	if err := Config.DB.Where("id = ?", id).Delete(user).Error; err != nil{
		return err
	}
	return nil
}