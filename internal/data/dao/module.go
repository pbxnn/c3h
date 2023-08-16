package dao

import "time"

const ModuleInfoTblName = "c3h_module_info"

type ModuleInfo struct {
	ID        uint
	Name      string
	Key       string
	Desc      string
	CreatedAt time.Time
	UpdatedAy time.Time
}

func (m *ModuleInfo) TableName() string {
	return ModuleDataMapTblName
}
