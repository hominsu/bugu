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

	"github.com/hominsu/bugu/app/bugu/service/internal/data/ent/artifact"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	buguV1 "github.com/hominsu/bugu/api/bugu/service/v1"
	"github.com/hominsu/bugu/app/bugu/service/internal/biz"
	"github.com/hominsu/bugu/app/bugu/service/internal/data/ent"
	"github.com/hominsu/bugu/app/bugu/service/internal/data/ent/file"
	"github.com/hominsu/bugu/app/bugu/service/internal/data/ent/user"
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

func (r *fileRepo) CreateFileMetadata(ctx context.Context, userId uuid.UUID, f *biz.File) (*biz.File, bool, error) {
	ok := false
	po, err := r.data.db.File.Create().
		SetID(f.ID).
		SetFileSha1(f.FileSha1).
		SetFileSize(f.FileSize).
		SetFileAddr(f.FileAddr).
		AddAffiliatedUserIDs(userId).
		Save(ctx)
	if err != nil && ent.IsConstraintError(err) {
		err = r.data.db.File.Update().
			Where(file.FileSha1EQ(f.FileSha1)).
			AddAffiliatedUserIDs(userId).
			Exec(ctx)
		if err != nil && ent.IsConstraintError(err) {
			return nil, false, buguV1.ErrorCreateConflict("create conflict, err: %v", err)
		}
		po, err = r.data.db.File.Query().
			Where(file.FileSha1EQ(f.FileSha1)).
			Only(ctx)
		ok = true
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, false, buguV1.ErrorUnknownError("unknown err: %v", err)
	}

	return &biz.File{
		ID:       po.ID,
		FileSha1: po.FileSha1,
		FileSize: po.FileSize,
		FileAddr: po.FileAddr,
		Type:     &po.Type,
	}, ok, nil
}

func (r *fileRepo) UpdateFileMetadata(ctx context.Context, file *biz.File) (*biz.File, error) {
	po, err := r.data.db.File.UpdateOneID(file.ID).
		SetFileSize(file.FileSize).
		SetFileAddr(file.FileAddr).
		SetNillableType(file.Type).
		Save(ctx)
	if err != nil && ent.IsConstraintError(err) {
		return nil, buguV1.ErrorCreateConflict("update conflict, err: %v", err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, buguV1.ErrorUnknownError("unknown err: %v", err)
	}

	return &biz.File{
		ID:       po.ID,
		FileSha1: po.FileSha1,
		FileSize: po.FileSize,
		FileAddr: po.FileAddr,
		Type:     &po.Type,
	}, nil
}

func (r *fileRepo) AppendFileMetadataToUser(ctx context.Context, userId, fileId uuid.UUID) (*biz.File, error) {
	po, err := r.data.db.File.UpdateOneID(fileId).
		AddAffiliatedUserIDs(userId).
		Save(ctx)
	if err != nil && ent.IsConstraintError(err) {
		return nil, buguV1.ErrorCreateConflict("update conflict, err: %v", err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, buguV1.ErrorUnknownError("unknown err: %v", err)
	}

	return &biz.File{
		ID:       po.ID,
		FileSha1: po.FileSha1,
		FileSize: po.FileSize,
		FileAddr: po.FileAddr,
		Type:     &po.Type,
	}, nil
}

func (r *fileRepo) GetFileMetadata(ctx context.Context, userId, fileId uuid.UUID) (*biz.File, error) {
	target, err := r.data.db.File.Query().
		Where(file.And(
			file.HasAffiliatedUserWith(user.IDEQ(userId)),
			file.IDEQ(fileId),
		)).
		Only(ctx)
	if err != nil && ent.IsNotFound(err) {
		return nil, buguV1.ErrorNotFoundError("find fileId: %s not found, err: %v", fileId.String(), err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, buguV1.ErrorUnknownError("unknown err: %v", err)
	}

	return &biz.File{
		ID:       target.ID,
		FileSha1: target.FileSha1,
		FileSize: target.FileSize,
		FileAddr: target.FileAddr,
		Type:     &target.Type,
	}, nil
}

func (r *fileRepo) GetFileMetadataByUserId(ctx context.Context, userId uuid.UUID) ([]*biz.File, error) {
	targets, err := r.data.db.File.Query().
		Where(file.And(
			file.HasAffiliatedUserWith(user.IDEQ(userId)),
		)).
		All(ctx)
	if err != nil && ent.IsNotFound(err) {
		return nil, buguV1.ErrorNotFoundError("find userId: %s not found, err: %v", userId.String(), err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, buguV1.ErrorUnknownError("unknown err: %v", err)
	}

	var rets []*biz.File
	for _, target := range targets {
		rets = append(rets, &biz.File{
			ID:       target.ID,
			FileSha1: target.FileSha1,
			FileSize: target.FileSize,
			FileAddr: target.FileAddr,
			Type:     &target.Type,
		})
	}

	return rets, nil
}

func (r *fileRepo) DeleteFileMetadata(ctx context.Context, userId, fileId uuid.UUID) error {
	err := r.data.db.File.UpdateOneID(fileId).
		RemoveAffiliatedUserIDs(userId).
		Exec(ctx)
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return buguV1.ErrorUnknownError("unknown err: %v", err)
	}

	err = r.data.db.Artifact.Update().
		Where(artifact.And(
			artifact.AffiliatedFileID(fileId),
			artifact.HasAffiliatedUserWith(user.IDEQ(userId)),
		)).
		RemoveAffiliatedUserIDs(userId).
		Exec(ctx)
	if err != nil && !ent.IsNotFound(err) {
		r.log.Errorf("unknown err: %v", err)
		return buguV1.ErrorUnknownError("unknown err: %v", err)
	}

	return nil
}

func (r *fileRepo) QueryByFileSha1(ctx context.Context, sha1 string) (*biz.File, error) {
	po, err := r.data.db.File.Query().
		Where(file.FileSha1EQ(sha1)).
		Only(ctx)
	if err != nil && ent.IsNotFound(err) {
		return nil, buguV1.ErrorNotFoundError("find sha1: %s not found, err: %v", sha1, err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, buguV1.ErrorUnknownError("unknown err: %v", err)
	}

	return &biz.File{
		ID:       po.ID,
		FileSha1: po.FileSha1,
		FileSize: po.FileSize,
		FileAddr: po.FileAddr,
		Type:     &po.Type,
	}, nil
}
