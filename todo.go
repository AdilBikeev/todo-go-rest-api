package todo_app

type TodoList struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
}

type UsersList struct {
	Id int `json:"id"`
	UserId int `json:"title"`
	ListId int `json:"description"`
}

type TodoItem struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Done string `json:"done"`
}

type ListsItem struct {
	Id int
	UserId int
	ListId int
}
