package service

import (
	"github.com/google/wire"
)

// SvcProviderSet is service providers.
var SvcProviderSet = wire.NewSet(
	NewProductNetService,
	NewControlNetService,
	NewR401SService,
)
