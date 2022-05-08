package data

import (
	"os"
	"testing"
	"time"

	"bugu/app/bugu/service/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/durationpb"
)

var (
	c = &conf.Data{
		Database: &conf.Data_Database{
			Driver: "mysql",
			Source: "root:dangerous@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local",
		},
		Redis: &conf.Data_Redis{
			Addr:            "127.0.0.1:6379",
			Db:              1,
			CacheExpiration: durationpb.New(time.Second * 1800),
			ReadTimeout:     durationpb.New(time.Millisecond * 200),
			WriteTimeout:    durationpb.New(time.Millisecond * 200),
		},
	}

	data *Data
)

func TestNewData(t *testing.T) {
}

func TestMain(m *testing.M) {
	var err error
	var cleanup func()

	logger := log.With(log.NewStdLogger(os.Stdout))
	helper := log.NewHelper(logger)

	entClient := NewEntClient(c, logger)
	redisCmd := NewRedisCmd(c, logger)

	data, cleanup, err = NewData(entClient, redisCmd, c, logger)
	if err != nil {
		helper.Fatal(err)
	}
	defer cleanup()

	ret := m.Run()
	os.Exit(ret)
}
