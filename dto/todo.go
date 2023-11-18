package dto

type NewTodosRequest struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type NewTodosResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Result     string `json:"result"`
}

type TodoWithId struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type ReadDataTodoResponse struct {
	StatusCode int          `json:"status_code"`
	Message    string       `json:"message"`
	Todos      []TodoWithId `json:"data"`
}
type ReadDataByIdTodoResponse struct {
	StatusCode int        `json:"status_code"`
	Message    string     `json:"message"`
	Todo       TodoWithId `json:"data"`
}
