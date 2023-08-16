package biz

import (
	"c3h/internal/data/dao"
	"context"
)

type CollectRepo interface {
	CollectData(ctx context.Context) ([]*dao.DataInfo, error)
	UpdateDB(ctx context.Context, list []*dao.DataInfo) error
	UpdateCache(ctx context.Context, list []*dao.DataInfo) error
}
