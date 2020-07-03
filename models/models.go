package models

type User struct {
	ID        string `json:"id,omitempty"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
	Pin       string `json:"pin"`
}

type Post struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Response struct {
	Status string `json:"status"`
	Error  bool   `json:"error"`
	Msg    string `json:"msg,omitempty"`
	Token  string `json:"token,omitempty"`
}
