package main


type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task" validate:"required,min=3"`
}

var todos = []Todo{
	{ID: 1, Task: "Learn Validation"},
	{ID: 1, Task: "Add Middleware"},
}



func main{
	
}