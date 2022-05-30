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

	buguV1 "bugu/api/bugu/service/v1"
	"bugu/app/bugu/service/internal/biz"
	"bugu/app/bugu/service/internal/data/ent"

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

func (r *artifactRepo) CreateArtifactMetadata(ctx context.Context, artifact *biz.Artifact) (*biz.Artifact, error) {
	po, err := r.data.db.Artifact.Create().
		SetAffiliatedFileID(artifact.FileID).
		SetID(artifact.ID).
		SetMethod(artifact.Method).
		Save(ctx)
	if err != nil && ent.IsConstraintError(err) {
		return nil, buguV1.ErrorCreateConflict("create conflict, err: %v", err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, buguV1.ErrorUnknownError("unknown err: %v", err)
	}

	return &biz.Artifact{
		ID:     po.ID,
		FileID: po.FileID,
		Method: po.Method,
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
		ID:     po.ID,
		FileID: po.FileID,
		Method: po.Method,
	}, nil
}

func (r *artifactRepo) GetArtifactMetadata(ctx context.Context, id uuid.UUID) (*biz.Artifact, error) {
	target, err := r.data.db.Artifact.Get(ctx, id)
	if err != nil && ent.IsNotFound(err) {
		return nil, buguV1.ErrorNotFoundError("find file id: %s not found, err: %v", id.String(), err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, buguV1.ErrorUnknownError("unknown err: %v", err)
	}

	return &biz.Artifact{
		ID:     target.ID,
		FileID: target.FileID,
		Method: target.Method,
	}, nil
}
