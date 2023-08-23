package biz

import (
	"context"

	"c3h/internal/data/dao"
)

type DataInfoRepo interface {
	GetByKeys(ctx context.Context, keys []string) ([]*dao.DataInfo, error)
	UpdateByKey(ctx context.Context, key string, up map[string]interface{}) (*dao.DataInfo, error)
	BatchUpdate(ctx context.Context, list []*dao.DataInfo) error
}
