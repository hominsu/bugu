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

	obfusionV1 "bugu/api/obfusion/service/v1"
	"bugu/app/bugu/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.ObfusionRepo = (*obfusionRepo)(nil)

type obfusionRepo struct {
	data *Data
	log  *log.Helper
}

// NewObfusionRepo .
func NewObfusionRepo(data *Data, logger log.Logger) biz.ObfusionRepo {
	return &obfusionRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/file")),
	}
}

func (r *obfusionRepo) Obfusion(ctx context.Context, data *biz.Obfusion) (*biz.Obfusion, error) {
	reply, err := r.data.oc.Obfusion(ctx, &obfusionV1.ObfusionRequest{
		Data: [][]byte{data.Data},
		Size: data.Size,
	})
	if err != nil {
		return nil, err
	}

	rd := []byte{}
	for _, datum := range reply.Data {
		rd = append(rd, datum...)
	}

	return &biz.Obfusion{
		Data: rd,
		Size: reply.Size,
	}, nil
}
