package data

import (
	"context"
	"time"

	"bugu/app/bugu/service/internal/conf"
	"bugu/app/bugu/service/internal/data/ent"
	"bugu/app/bugu/service/internal/data/ent/migrate"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewEntClient, NewRedisCmd, NewUserRepo, NewFileRepo, NewArtifactRepo)

// Data .
type Data struct {
	db    *ent.Client
	rdCmd redis.Cmdable

	conf *conf.Data
}

// NewData .
func NewData(entClient *ent.Client, rdCmd redis.Cmdable, conf *conf.Data, logger log.Logger) (*Data, func(), error) {
	helper := log.NewHelper(log.With(logger, "module", "ecode-service/data"))

	d := &Data{
		db:    entClient,
		rdCmd: rdCmd,
		conf:  conf,
	}
	return d, func() {
		if err := d.db.Close(); err != nil {
			helper.Error(err)
		}
	}, nil
}

func NewEntClient(conf *conf.Data, logger log.Logger) *ent.Client {
	helper := log.NewHelper(log.With(logger, "module", "ecode-service/data/ent"))

	client, err := ent.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	if err != nil {
		helper.Fatalf("failed opening connection to db: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
		helper.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

func NewRedisCmd(conf *conf.Data, logger log.Logger) redis.Cmdable {
	helper := log.NewHelper(log.With(logger, "module", "ecode-service/data/redis"))

	client := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		DB:           int(conf.Redis.Db),
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		DialTimeout:  time.Second * 2,
		PoolSize:     10,
	})
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFunc()
	err := client.Ping(timeout).Err()
	if err != nil {
		helper.Fatalf("redis connect error: %v", err)
	}
	return client
}
