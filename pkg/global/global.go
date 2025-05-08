package global

import (
	"user/pkg/config"
	"user/pkg/types"

	"github.com/dubbogo/gost/log/logger"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	ConfigCenterConfig *config.ConfigCenterConfig

	Config *config.Config
	DB     *gorm.DB
	Redis  redis.UniversalClient
	Log    func() logger.Logger
	Cache  types.Cache
)
