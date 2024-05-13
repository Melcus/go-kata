package user

type user struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type publicUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
