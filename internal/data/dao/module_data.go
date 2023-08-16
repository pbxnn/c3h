package dao

import "time"

const ModuleDataMapTblName = "c3h_module_data_map"

type ModuleDataMap struct {
	ID        uint       `gorm:"primaryKey;column:id"`
	DataKey   string     `gorm:"column:data_key"`
	ModuleKey string     `gorm:"column:module_key"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func (m *ModuleDataMap) TableName() string {
	return ModuleDataMapTblName
}
