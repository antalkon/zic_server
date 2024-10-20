package models

type NewDataTg struct {
	Token    string `json:"token"`
	SysStat  bool   `json:"sysstat`
	NewTasks bool   `json:"newtasks"`
}
