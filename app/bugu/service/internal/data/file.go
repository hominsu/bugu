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

	v1 "bugu/api/bugu/service/v1"
	"bugu/app/bugu/service/internal/biz"
	"bugu/app/bugu/service/internal/data/ent"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

var _ biz.FileRepo = (*fileRepo)(nil)

type fileRepo struct {
	data *Data
	log  *log.Helper
}

// NewFileRepo .
func NewFileRepo(data *Data, logger log.Logger) biz.FileRepo {
	return &fileRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/file")),
	}
}

func (r *fileRepo) CreateFileMetadata(ctx context.Context, file *biz.File) (*biz.File, error) {
	//md, ok := metadata.FromClientContext(ctx)
	//if !ok {
	//	return nil, v1.ErrorInternalServerError("Openid does not exist in context")
	//}
	//userid := md.Get("x-md-global-userid")

	po, err := r.data.db.File.Create().
		SetID(file.ID).
		SetFileSha1(file.FileSha1).
		SetFileSize(file.FileSize).
		SetFileAddr(file.FileAddr).
		Save(ctx)
	if err != nil && ent.IsConstraintError(err) {
		return nil, v1.ErrorCreateConflict("create conflict, err: %v", err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, v1.ErrorUnknownError("unknown err: %v", err)
	}

	return &biz.File{
		ID:       po.ID,
		FileSha1: po.FileSha1,
		FileSize: po.FileSize,
		FileAddr: po.FileAddr,
	}, nil
}

func (r *fileRepo) UpdateFileMetadata(ctx context.Context, file *biz.File) (*biz.File, error) {
	po, err := r.data.db.File.UpdateOneID(file.ID).
		SetFileSize(file.FileSize).
		SetFileAddr(file.FileAddr).
		SetNillableType(file.Type).
		Save(ctx)
	if err != nil && ent.IsConstraintError(err) {
		return nil, v1.ErrorCreateConflict("update conflict, err: %v", err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, v1.ErrorUnknownError("unknown err: %v", err)
	}

	return &biz.File{
		ID:       po.ID,
		FileSha1: po.FileSha1,
		FileSize: po.FileSize,
		FileAddr: po.FileAddr,
	}, nil
}

func (r *fileRepo) GetFileMetadata(ctx context.Context, id uuid.UUID) (*biz.File, error) {
	target, err := r.data.db.File.Get(ctx, id)
	if err != nil && ent.IsNotFound(err) {
		return nil, v1.ErrorNotFoundError("find file id: %s not found, err: %v", id.String(), err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, v1.ErrorUnknownError("unknown err: %v", err)
	}

	return &biz.File{
		ID:       target.ID,
		FileSha1: target.FileSha1,
		FileSize: target.FileSize,
		FileAddr: target.FileAddr,
	}, nil
}
