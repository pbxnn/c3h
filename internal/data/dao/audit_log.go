package dao

import (
	"time"
)

const AuditLogTblName = "c3h_audit_log"

type AuditLog struct {
	ID        uint64 `gorm:"primary"`
	UserID    uint64
	Operation uint
	Resource  uint
	IsSuccess uint
	Details   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (al *AuditLog) TableName() string {
	return AuditLogTblName
}
