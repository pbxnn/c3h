package biz

import (
	"context"

	"c3h/internal/data/dao"
)

type ModuleRelationRepo interface {
	UpdateModuleRelationCache(ctx context.Context)
	GetByModuleKey(ctx context.Context, moduleKey string) ([]*dao.ModuleDataMap, error)
}
