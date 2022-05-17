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
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

// Artifact is the model entity for the Artifact schema.
type Artifact struct {
	ID           uuid.UUID `json:"id,omitempty"`
	FileID       uuid.UUID `json:"file_id,omitempty"`
	ArtifactHash uuid.UUID `json:"artifact_hash,omitempty"`
	ArtifactSize string    `json:"artifact_size,omitempty"`
	ArtifactAddr string    `json:"artifact_addr,omitempty"`
}

type ArtifactRepo interface{}

type ArtifactUsecase struct {
	repo ArtifactRepo

	log *log.Helper
}

func NewArtifactUsecase(repo ArtifactRepo, logger log.Logger) *ArtifactUsecase {
	return &ArtifactUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}
