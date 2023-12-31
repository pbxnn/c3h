package data

import (
	"context"

	"c3h/api/platform"
	"c3h/internal/conf"
	"c3h/pkg/cache"

	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// DProviderSet is data providers.
var DProviderSet = wire.NewSet(
	NewData,
	NewDB,
	NewCache,
	//NewIndustrialDataClient,

	NewControlNetRepo,
	NewProductNetRepo,
	NewAuditLogRepo,
	NewCollectorRepo,
	NewModuleRelationRepo,
	NewR401SRepo,
	NewDataInfoRepo,
	NewPlatformClient,
)

const defaultCacheSize = 1024000000 // 100MB
const defaultCacheExpire = 3        // 3s

// Data .
type Data struct {
	log   *log.Helper
	db    *gorm.DB
	cache cache.Cache
}

// NewData .
func NewData(c *conf.Data, db *gorm.DB, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		db: db,
	}, cleanup, nil
}

func NewDB(c *conf.Data, logger log.Logger) *gorm.DB {
	log := log.NewHelper(logger)
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return db
	}

	//log.Infof("init DB...")
	//err = db.Use(&gp.OpentracingPlugin{})
	//if err != nil {
	//	log.Fatal(err)
	//	return db
	//}

	return db
}

func NewPlatformClient(c *conf.Data, logger log.Logger) platform.PlatformHTTPClient {
	conn, err := http.NewClient(
		context.Background(),
		http.WithEndpoint(c.IndustrialApi.Addr),
		http.WithMiddleware(
			recovery.Recovery(),
			logging.Client(logger),
		),
	)

	if err != nil {
		panic(err)
	}
	client := platform.NewPlatformHTTPClient(conn)
	return client
}

func NewCache(c *conf.Data) cache.Cache {
	size := int64(defaultCacheSize)
	expire := int64(defaultCacheExpire)
	if cc := c.GetCache(); cc != nil {
		if s := cc.GetSize(); s > 0 {
			size = (int64)(s)
		}

		if e := cc.GetExpire(); e > 0 {
			expire = (int64)(e)
		}
	}

	return cache.NewCache(size, expire)
}
