package models

type SystemConfig struct {
	BaseModel
	Key   string `json:"key"`
	Value string `json:"value"`
}
