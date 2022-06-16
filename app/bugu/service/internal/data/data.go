/*
 * MIT License
 *
 * Copyright (c) 2022. HominSu
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 */

package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	obfusionV1 "github.com/hominsu/bugu/api/obfusion/service/v1"
	packerV1 "github.com/hominsu/bugu/api/packer/service/v1"
	"github.com/hominsu/bugu/app/bugu/service/internal/conf"
	"github.com/hominsu/bugu/app/bugu/service/internal/data/ent"
	"github.com/hominsu/bugu/app/bugu/service/internal/data/ent/migrate"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewEntClient,
	NewRedisCmd,
	NewOubfusionServiceClient,
	NewUserRepo,
	NewFileRepo,
	NewArtifactRepo,
	NewObfusionRepo,
)

// Data .
type Data struct {
	db    *ent.Client
	rdCmd redis.Cmdable

	oc obfusionV1.BuguObfusionClient
	pc packerV1.BuguPackerClient

	conf *conf.Data
}

// NewData .
func NewData(entClient *ent.Client,
	rdCmd redis.Cmdable,
	oc obfusionV1.BuguObfusionClient,
	conf *conf.Data,
	logger log.Logger,
) (*Data, func(), error) {
	helper := log.NewHelper(log.With(logger, "module", "ecode-service/data"))

	d := &Data{
		db:    entClient,
		rdCmd: rdCmd,
		oc:    oc,
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

func NewOubfusionServiceClient(conf *conf.Server) obfusionV1.BuguObfusionClient {
	opts := []grpc.ClientOption{
		grpc.WithEndpoint("bugu-obfusion-service:9000"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	}
	if conf.Http.Timeout != nil {
		opts = append(opts, grpc.WithTimeout(conf.Http.Timeout.AsDuration()))
	}

	conn, err := grpc.DialInsecure(context.Background(), opts...)
	if err != nil {
		panic(err)
	}

	c := obfusionV1.NewBuguObfusionClient(conn)
	return c
}

func NewPackerServiceClient(conf *conf.Server) packerV1.BuguPackerClient {
	opts := []grpc.ClientOption{
		grpc.WithEndpoint("bugu-packer-service:9000"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	}
	if conf.Http.Timeout != nil {
		opts = append(opts, grpc.WithTimeout(conf.Http.Timeout.AsDuration()))
	}

	conn, err := grpc.DialInsecure(context.Background(), opts...)
	if err != nil {
		panic(err)
	}

	c := packerV1.NewBuguPackerClient(conn)
	return c
}
