package biz

import (
	"context"

	"c3h/internal/data/dao"
)

type ModuleRelationRepo interface {
	RefreshCache(ctx context.Context)
	GetByModuleKey(ctx context.Context, moduleKey string) ([]*dao.ModuleDataMap, error)
}
