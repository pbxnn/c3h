package biz

import "github.com/google/wire"

// BProviderSet is biz providers.
var BProviderSet = wire.NewSet(
	NewProductNetUsecase,
	NewControlNetUsecase,
	NewAuditLogUsecase,
)
