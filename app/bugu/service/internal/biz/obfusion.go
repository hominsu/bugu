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

type Obfusion struct {
	Data []byte `json:"data,omitempty"`
	Size uint32 `json:"size,omitempty"`
}

type ObfusionRepo interface {
	Obfusion(ctx context.Context, data *Obfusion) (*Obfusion, error)
}

type ObfusionUsecase struct {
	repo ObfusionRepo

	log *log.Helper
}

func NewObfusionUsecase(repo ObfusionRepo, logger log.Logger) *ObfusionUsecase {
	return &ObfusionUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *ObfusionUsecase) Obfusion(ctx context.Context, data *Obfusion) (*Obfusion, error) {
	return uc.repo.Obfusion(ctx, data)
}
