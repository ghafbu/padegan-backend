package bootstrap

import (
	"github.com/ghafbu/padegan-backend/app/model"
	"github.com/ghafbu/padegan-backend/app/seeder"
	"github.com/ghafbu/padegan-backend/config"
	"github.com/ghafbu/padegan-backend/lib/consts"
	_interface "github.com/ghafbu/padegan-backend/lib/interface"
	"github.com/ghafbu/padegan-backend/lib/interfaceimplement/giiilogger"
	"github.com/ghafbu/padegan-backend/lib/interfaceimplement/giiiratelimiterstorage"
	"github.com/ghafbu/padegan-backend/lib/interfaceimplement/giiiredis"
	"github.com/ghafbu/padegan-backend/lib/interfaceimplement/giiitoken"
	messages "github.com/ghafbu/padegan-backend/lib/message"
	"github.com/ghafbu/padegan-backend/pkg/databasepkg/gorm"
	"github.com/ghafbu/padegan-backend/pkg/databasepkg/gorm/migrator"
	"github.com/ghafbu/padegan-backend/pkg/databasepkg/gorm/postgres"
)

type Application struct {
	Config             config.Config
	Logger             _interface.ILogger
	Token              _interface.IToken
	Redis              _interface.IRedis
	RateLimiterStorage _interface.IRateLimiterStorage
	Postgres           gorm.DB
}

func Start(paramConfigPath string) (*Application, error) {
	var (
		thisApp = &Application{}
		err     error
	)

	// Init Config
	if thisApp.Config, err = config.Load(paramConfigPath); err != nil {
		return nil, err
	}

	// Init Logger
	thisApp.Logger = giiilogger.NewILogger(
		giiilogger.ImplementNameZap,
		thisApp.Config.Logger,
	)

	// Init Token
	thisApp.Token = giiitoken.NewIToken(giiitoken.ImplementNameJWT)

	// Init Redis
	thisApp.Redis, err = giiiredis.NewIRedis(giiiredis.ImplementNameRedis, thisApp.Config.Redis)
	if err != nil {
		return nil, err
	}

	// Init RateLimiter Storage
	if thisApp.RateLimiterStorage, err = giiiratelimiterstorage.NewIRedis(
		giiiratelimiterstorage.ImplementNameRedis,
		thisApp.Config.Redis,
		consts.RateLimiterStorageTimeout,
	); err != nil {
		return nil, err
	}

	// Init Postgres
	thisApp.Postgres, err = postgres.Init(thisApp.Config.Postgres)
	if err != nil {
		return nil, err
	}

	// Init Migrations
	migrator.Init(thisApp.Postgres).
		Down().
		LoadEntities(&model.User{}, &model.UserSetting{}).
		AutoMigrate()

	// Init Seeders
	err = seeder.UsersSeeders(100, thisApp.Postgres)
	if err != nil {
		return nil, err
	}

	// Init Messages
	messages.Init()

	return thisApp, nil
}
