package interfaces

type RequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
