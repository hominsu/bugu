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

	"github.com/go-kratos/kratos/v2/log"
	packerV1 "github.com/hominsu/bugu/api/packer/service/v1"
	"github.com/hominsu/bugu/app/bugu/service/internal/biz"
)

var _ biz.PackerRepo = (*packerRepo)(nil)

type packerRepo struct {
	data *Data
	log  *log.Helper
}

// NewPackerRepo .
func NewPackerRepo(data *Data, logger log.Logger) biz.PackerRepo {
	return &packerRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/file")),
	}
}

func (r *packerRepo) Packer(ctx context.Context, data *biz.Packer) (*biz.Packer, error) {
	reply, err := r.data.pc.Packer(ctx, &packerV1.PackerRequest{
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

	return &biz.Packer{
		Data: rd,
		Size: reply.Size,
	}, nil
}
