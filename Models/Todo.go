package Models

import (
	"time"
	"goapp/Config"
	_ "github.com/go-sql-driver/mysql"

)

type ToDo struct {
	Id        uint		`json:"id"`
	UserId    uint		`json:"user"`
	Task      string	`json:"task"`
	Status 	  bool		`json:"status"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (b *ToDo) TableName() string {
	return "todos"
}


func CreateATodo(todo *ToDo) (err error) {
	if err = Config.DB.Create(todo).Error; err != nil{
		return err
	}
	return nil
}

func GetTodos(todos *[]ToDo, userId string, limit string, offset string, search string) (err error){
	chain := Config.DB.Where("user_id = ? ", userId)
	if search != ""{
		chain = chain.Where("task LIKE ?", "%"+search+"%")
	}
	if err = chain.Limit(limit).Offset(offset).Find(todos).Error; err != nil {
		return err
	}
	return nil
}

func GetATodo(todo *ToDo, id string, userId string) error{
	if err := Config.DB.Where("id = ? AND user_id = ?", id, userId).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func UpdateATodo(todo *ToDo)(err error){
	if err = Config.DB.Save(todo).Error; err != nil{
		return err
	}
	return nil
}
func DeleteATodo(todo *ToDo, id string, userId string)(err error){
	if err = Config.DB.Where("id = ?", id).Delete(todo).Error; err != nil{
		return err
	}
	return nil
}