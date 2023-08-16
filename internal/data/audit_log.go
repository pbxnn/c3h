package data

import (
	"c3h/internal/biz"
	"c3h/internal/data/dao"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

var _ biz.AuditLogRepo = (*auditLogRepo)(nil)

type auditLogRepo struct {
	logger *log.Helper
	db     *gorm.DB
}

func NewAuditLogRepo(db *gorm.DB, logger log.Logger) biz.AuditLogRepo {
	return &auditLogRepo{
		db:     db,
		logger: log.NewHelper(log.With(logger, "module", "audit-log-repo")),
	}
}

func (alr *auditLogRepo) AddRecord(ctx context.Context, d *dao.AuditLog) error {
	return alr.db.Create(d).Error
}

func (alr *auditLogRepo) ListRecord(ctx context.Context, conds map[string]interface{}) ([]*dao.AuditLog, error) {

	var res []*dao.AuditLog
	err := alr.db.WithContext(ctx).Where(conds).Find(&res).Error
	return res, err
}
