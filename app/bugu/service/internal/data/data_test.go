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
	"os"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/hominsu/bugu/app/bugu/service/internal/conf"
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

	data, cleanup, err = NewData(entClient, redisCmd, nil, c, logger)
	if err != nil {
		helper.Fatal(err)
	}
	defer cleanup()

	ret := m.Run()
	os.Exit(ret)
}
