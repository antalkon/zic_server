package models

type SActivate struct {
	User     string `json:"user`
	Password string `json:"password`
	Port     string `json:"port"`
	Code     string `json:"code"`
}

type SActivateRequest struct {
	PIP      string `json:"serverIP"`
	LIP      string `json:"serverLocalIP"`
	Version  string `json:"serverVersion"`
	User     string `json:"user"`
	Password string `json:"password"`
	Port     string `json:"port"`
}
