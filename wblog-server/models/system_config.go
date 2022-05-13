package models

type SystemConfig struct {
	BaseModel
	Key   string `json:"key"  gorm:"type:varchar(64);"`
	Value string `json:"value" gorm:"type:varchar(64);"`
}
