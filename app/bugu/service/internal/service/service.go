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

package service

import (
	buguV1 "bugu/api/bugu/service/v1"
	"bugu/app/bugu/service/internal/biz"
	"bugu/app/bugu/service/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewBuguService, NewBuguFileService)

type BuguService struct {
	buguV1.UnimplementedBuguServer

	uu  *biz.UserUsecase
	au  *biz.ArtifactUsecase
	log *log.Helper
}

func NewBuguService(uu *biz.UserUsecase, logger log.Logger) *BuguService {
	return &BuguService{
		uu:  uu,
		log: log.NewHelper(log.With(logger, "module", "service/bugu")),
	}
}

type BuguFileService struct {
	buguV1.UnimplementedBuguFileServer

	fu  *biz.FileUsecase
	dc  *conf.Data
	log *log.Helper
}

func NewBuguFileService(fu *biz.FileUsecase, dc *conf.Data, logger log.Logger) *BuguFileService {
	return &BuguFileService{
		fu:  fu,
		dc:  dc,
		log: log.NewHelper(log.With(logger, "module", "service/bugu-file")),
	}
}
