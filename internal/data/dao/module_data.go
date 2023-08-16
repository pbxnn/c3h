package dao

import "time"

const ModuleDataMapTblName = "c3h_module_data_map"

type ModuleDataMap struct {
	ID        uint
	DataID    uint
	ModuleID  uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *ModuleDataMap) TableName() string {
	return ModuleDataMapTblName
}
