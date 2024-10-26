package models

type NewDataTg struct {
	Token    string `json:"token"`
	SysStat  bool   `json:"sysSt`
	NewTasks bool   `json:"newTask"`
}

type GlNetwork struct {
	Global bool `json:"global"`
}

type SecFirewall struct {
	Firewall bool `json:"firewall"`
}

type ReqAuth struct {
	Auth bool `json:"auth"`
}
