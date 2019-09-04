package shared

type UserService interface {
	GetByName(n string) (*UserServiceResponse, error)
}

type UserServiceResponse struct {
	Valid   bool   `json:"valid"`
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
}
