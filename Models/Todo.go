package Models

import (
	"time"
	"goapp/Config"
	_ "github.com/go-sql-driver/mysql"

)

type ToDo struct {
	Id        uint		`json:"id"`
	UserId    uint		`json:"user"`
	Task      string	`json:"task" validate:"required,max=255,min=3"`
	Status 	  bool		`json:"status"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (b *ToDo) TableName() string {
	return "todos"
}


func CreateATodo(todo *ToDo) error {
	if err := Config.DB.Create(todo).Error; err != nil{
		return err
	}
	return nil
}

func GetTodos(todos *[]ToDo, userId string, limit int, offset int, search string) (err error){
	chain := Config.DB.Where("user_id = ? ", userId)
	if search != ""{
		chain = chain.Where("task LIKE ?", "%"+search+"%")
	}
	if err = chain.Limit(limit).Offset(offset).Find(todos).Error; err != nil {
		return err
	}
	return nil
}


func CountTodos(userId string, search string) (count int64, err error){
	chain := Config.DB.Model(&ToDo{}).Where("user_id = ? ", userId)
	if search != ""{
		chain = chain.Where("task LIKE ?", "%"+search+"%")
	}
	if err = chain.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func GetATodo(todo *ToDo, id string, userId string) error{
	if err := Config.DB.Where("id = ? AND user_id = ?", id, userId).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func UpdateATodo(todo *ToDo) error {
	if err := Config.DB.Save(todo).Error; err != nil{
		return err
	}
	return nil
}
func DeleteATodo(todo *ToDo, id string, userId string) error{
	if err := Config.DB.Where("id = ?", id).Delete(todo).Error; err != nil{
		return err
	}
	return nil
}