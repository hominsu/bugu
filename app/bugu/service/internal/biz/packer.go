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

package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Packer struct {
	Data []byte `json:"data,omitempty"`
	Size uint32 `json:"size,omitempty"`
}

type PackerRepo interface {
	Packer(ctx context.Context, data *Packer) (*Packer, error)
}

type PackerUsecase struct {
	repo PackerRepo

	log *log.Helper
}

func NewPackerUsecase(repo PackerRepo, loggger log.Logger) *PackerUsecase {
	return &PackerUsecase{
		repo: repo,
		log:  log.NewHelper(loggger),
	}
}

func (uc *PackerUsecase) Packer(ctx context.Context, data *Packer) (*Packer, error) {
	return uc.repo.Packer(ctx, data)
}
