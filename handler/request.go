package handler

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   string `json:"remote"`
	Link     string `json:"link"`
	Salary   string `json:"salary"`
}
