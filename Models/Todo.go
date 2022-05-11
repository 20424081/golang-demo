package Models

import (
	"time"
	"goapp/Config"
)

type ToDo struct {
	Id        uint		`json:"id"`
	UserId    uint		`json:"user"`
	Task      string	`json:"task"`
	Status 	  bool		`json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateATodo(todo *Todo) (err error) {
}

func GetTodos()(err error){

}

func GetATodo(todo *Todo, id string)(err error){

}

func UpdateATodo(todo *Todo, id string)(err error){

}
func DeleteATodo(todo *Todo, id string)(err error){

}