package dao

import "time"

const DataInfoTblName = "c3h_module_data_map"

type DataInfo struct {
	ID        uint
	Name      string
	Key       string
	Desc      string
	Type      uint
	RealValue float64
	SetValue  float64
	LowLimit  float64
	HighLimit float64
	Unit      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (d *DataInfo) TableName() string {
	return DataInfoTblName
}
