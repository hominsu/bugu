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

	buguV1 "github.com/hominsu/bugu/api/bugu/service/v1"
	"github.com/hominsu/bugu/app/bugu/service/internal/biz"
	"github.com/hominsu/bugu/app/bugu/service/internal/data/ent"
	"github.com/hominsu/bugu/app/bugu/service/internal/data/ent/artifact"
	"github.com/hominsu/bugu/app/bugu/service/internal/data/ent/file"
	"github.com/hominsu/bugu/app/bugu/service/internal/data/ent/user"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

var _ biz.ArtifactRepo = (*artifactRepo)(nil)

type artifactRepo struct {
	data *Data
	log  *log.Helper
}

// NewArtifactRepo .
func NewArtifactRepo(data *Data, logger log.Logger) biz.ArtifactRepo {
	return &artifactRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/artifact")),
	}
}

func (r *artifactRepo) CreateArtifactMetadata(ctx context.Context, userId uuid.UUID, a *biz.Artifact) (*biz.Artifact, error) {
	po, err := r.data.db.Artifact.Create().
		SetAffiliatedFileID(a.AffiliatedFileID).
		SetFileID(a.FileID).
		SetID(a.ID).
		SetMethod(a.Method).
		AddAffiliatedUserIDs(userId).
		Save(ctx)
	if err != nil && ent.IsConstraintError(err) {
		return nil, buguV1.ErrorCreateConflict("create conflict, err: %v", err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, buguV1.ErrorUnknownError("unknown err: %v", err)
	}

	return &biz.Artifact{
		ID:               po.ID,
		FileID:           po.FileID,
		AffiliatedFileID: po.AffiliatedFileID,
		Method:           po.Method,
	}, nil
}

func (r *artifactRepo) UpdateArtifactMetadata(ctx context.Context, artifact *biz.Artifact) (*biz.Artifact, error) {
	po, err := r.data.db.Artifact.UpdateOneID(artifact.ID).
		SetMethod(artifact.Method).
		Save(ctx)
	if err != nil && ent.IsConstraintError(err) {
		return nil, buguV1.ErrorCreateConflict("update conflict, err: %v", err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, buguV1.ErrorUnknownError("unknown err: %v", err)
	}

	return &biz.Artifact{
		ID:               po.ID,
		FileID:           po.FileID,
		AffiliatedFileID: po.AffiliatedFileID,
		Method:           po.Method,
	}, nil
}

func (r *artifactRepo) AppendArtifactMetadataByArtifactFileIdToUser(ctx context.Context, userId, fileId uuid.UUID) (*biz.Artifact, error) {
	err := r.data.db.Artifact.Update().
		Where(artifact.FileIDEQ(fileId)).
		AddAffiliatedUserIDs(userId).
		Exec(ctx)
	if err != nil && ent.IsConstraintError(err) {
		return nil, buguV1.ErrorCreateConflict("update conflict, err: %v", err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, buguV1.ErrorUnknownError("unknown err: %v", err)
	}

	target, err := r.data.db.Artifact.Query().
		Where(artifact.FileIDEQ(fileId)).
		Only(ctx)
	if err != nil && ent.IsNotFound(err) {
		return nil, buguV1.ErrorNotFoundError("find fileId: %s not found, err: %v", fileId.String(), err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, buguV1.ErrorUnknownError("unknown err: %v", err)
	}

	return &biz.Artifact{
		ID:               target.ID,
		FileID:           target.FileID,
		AffiliatedFileID: target.AffiliatedFileID,
		Method:           target.Method,
	}, nil
}

func (r *artifactRepo) AppendArtifactMetadataToUser(ctx context.Context, userId, artifactId uuid.UUID) (*biz.Artifact, error) {
	po, err := r.data.db.Artifact.UpdateOneID(artifactId).
		AddAffiliatedUserIDs(userId).
		Save(ctx)
	if err != nil && ent.IsConstraintError(err) {
		return nil, buguV1.ErrorCreateConflict("update conflict, err: %v", err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, buguV1.ErrorUnknownError("unknown err: %v", err)
	}

	return &biz.Artifact{
		ID:               po.ID,
		FileID:           po.FileID,
		AffiliatedFileID: po.AffiliatedFileID,
		Method:           po.Method,
	}, nil
}

func (r *artifactRepo) GetArtifactMetadata(ctx context.Context, userId, artifactId uuid.UUID) (*biz.Artifact, error) {
	target, err := r.data.db.Artifact.Query().
		Where(artifact.And(
			artifact.HasAffiliatedUserWith(user.IDEQ(userId)),
			artifact.IDEQ(artifactId),
		)).
		Only(ctx)
	if err != nil && ent.IsNotFound(err) {
		return nil, buguV1.ErrorNotFoundError("find artifactId: %s not found, err: %v", artifactId.String(), err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, buguV1.ErrorUnknownError("unknown err: %v", err)
	}

	return &biz.Artifact{
		ID:               target.ID,
		FileID:           target.FileID,
		AffiliatedFileID: target.AffiliatedFileID,
		Method:           target.Method,
	}, nil
}

func (r *artifactRepo) GetArtifactMetadataByFileId(ctx context.Context, userId, fileId uuid.UUID) ([]*biz.Artifact, error) {
	targets, err := r.data.db.Artifact.Query().
		Where(artifact.And(
			artifact.HasAffiliatedUserWith(user.IDEQ(userId)),
			artifact.HasAffiliatedFileWith(file.IDEQ(fileId)),
		)).
		All(ctx)
	if err != nil && ent.IsNotFound(err) {
		return nil, buguV1.ErrorNotFoundError("find fileId: %s not found, err: %v", fileId.String(), err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, buguV1.ErrorUnknownError("unknown err: %v", err)
	}

	var rets []*biz.Artifact
	for _, target := range targets {
		rets = append(rets, &biz.Artifact{
			ID:               target.ID,
			FileID:           target.FileID,
			AffiliatedFileID: target.AffiliatedFileID,
			Method:           target.Method,
		})
	}

	return rets, nil
}

func (r *artifactRepo) DeleteArtifactMetadata(ctx context.Context, userId, artifactId uuid.UUID) error {
	po, err := r.data.db.Artifact.UpdateOneID(artifactId).
		RemoveAffiliatedUserIDs(userId).
		Save(ctx)
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return buguV1.ErrorUnknownError("unknown err: %v", err)
	}

	err = r.data.db.File.Update().
		Where(file.And(
			file.HasAffiliatedUserWith(user.IDEQ(userId)),
			file.IDEQ(po.FileID),
		)).
		RemoveAffiliatedUserIDs(userId).
		Exec(ctx)
	if err != nil && !ent.IsNotFound(err) {
		r.log.Errorf("unknown err: %v", err)
		return buguV1.ErrorUnknownError("unknown err: %v", err)
	}

	return nil
}
