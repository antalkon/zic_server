package models

// получение настроек тг
type NewDataTg struct {
	Token    string `json:"token"`
	SysStat  bool   `json:"sysSt`
	NewTasks bool   `json:"newTask"`
}

// настройки гл сети
type GlNetwork struct {
	Global bool `json:"global"`
}

// настройки фаирвала
type SecFirewall struct {
	Firewall bool `json:"firewall"`
}

// настройки запроса авторизации
type ReqAuth struct {
	Auth bool `json:"auth"`
}
