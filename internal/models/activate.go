package models

// приме запроса на активацию
type SActivate struct {
	User     string `json:"user`
	Password string `json:"password`
	Port     string `json:"port"`
	Code     string `json:"code"`
}

// запрос активации на сервера zic
type SActivateRequest struct {
	PIP      string `json:"serverIP"`
	LIP      string `json:"serverLocalIP"`
	Version  string `json:"serverVersion"`
	User     string `json:"user"`
	Password string `json:"password"`
	Port     string `json:"port"`
}
