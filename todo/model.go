package todo

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type TodoList struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type ToDoItem struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type UserList struct {
	UserID int
	ListID int
}

type ListItem struct {
	TodoID int
	ItemID int
}
