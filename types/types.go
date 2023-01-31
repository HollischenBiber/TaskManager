package types

type AuthData struct {
	Username string `json:"username"`
	Pass     string `json:"pass"`
}

type Answer struct {
	Success   bool   `json:"success"`
	ErrorText string `json:"errorText"`
}
