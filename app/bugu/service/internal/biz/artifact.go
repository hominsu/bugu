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
	"io/ioutil"
	"os"
	"path/filepath"

	buguV1 "github.com/hominsu/bugu/api/bugu/service/v1"
	"github.com/hominsu/bugu/app/bugu/service/internal/conf"
	"github.com/hominsu/bugu/pkg"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

// Artifact is the model entity for the Artifact schema.
type Artifact struct {
	ID     uuid.UUID `json:"id,omitempty"`
	FileID uuid.UUID `json:"file_id,omitempty"`
	Method string    `json:"method,omitempty"`
}

type ArtifactRepo interface {
	CreateArtifactMetadata(ctx context.Context, artifact *Artifact) (*Artifact, error)
	UpdateArtifactMetadata(ctx context.Context, artifact *Artifact) (*Artifact, error)
	GetArtifactMetadata(ctx context.Context, id uuid.UUID) (*Artifact, error)
}

type ArtifactUsecase struct {
	ar ArtifactRepo
	or ObfusionRepo
	fr FileRepo

	dc  *conf.Data
	log *log.Helper
}

func NewArtifactUsecase(ar ArtifactRepo, or ObfusionRepo, fr FileRepo, dc *conf.Data, logger log.Logger) *ArtifactUsecase {
	return &ArtifactUsecase{
		ar:  ar,
		or:  or,
		fr:  fr,
		dc:  dc,
		log: log.NewHelper(logger),
	}
}

func (uc *ArtifactUsecase) Confusion(ctx context.Context, fileId string) (*Artifact, error) {
	f, cleanup, err := uc.getFile(ctx, fileId)
	if err != nil {
		return nil, err
	}
	defer cleanup()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	dto, err := uc.or.Obfusion(ctx, &Obfusion{
		Data: data,
		Size: uint32(len(data)),
	})
	if err != nil {
		return nil, err
	}

	fileMetadata, err := uc.saveFile(ctx, dto, uc.dc.File.Path)
	if err != nil {
		return nil, err
	}

	u, err := uuid.NewRandom()
	if err != nil {
		return nil, buguV1.ErrorUuidGenerateFailed("create file uuid failed, err: %v", err)
	}

	artifactMetadata, err := uc.ar.CreateArtifactMetadata(ctx, &Artifact{
		ID:     u,
		FileID: fileMetadata.ID,
	})
	if err != nil {
		return nil, err
	}

	return &Artifact{
		ID:     artifactMetadata.ID,
		FileID: artifactMetadata.FileID,
		Method: artifactMetadata.Method,
	}, nil
}

func (uc *ArtifactUsecase) saveFile(ctx context.Context, oData *Obfusion, dir string) (*File, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return nil, buguV1.ErrorUuidGenerateFailed("create file uuid failed, err: %v", err)
	}

	dir = filepath.Join(dir, u.String())

	ok, err := pkg.PathExists(dir)
	if err != nil {
		return nil, buguV1.ErrorUnknownError("check file failed, err: %v", err)
	}
	if ok {
		return nil, buguV1.ErrorCreateConflict("create file conflict")
	}

	f, err := os.OpenFile(dir, os.O_RDWR|os.O_CREATE, 0o666)
	if err != nil {
		return nil, err
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			uc.log.Error(err)
		}
	}(f)

	metadata := &File{
		ID:       u,
		FileAddr: dir,
	}

	wSize, err := f.Write(oData.Data)
	if err != nil {
		return nil, err
	}

	metadata.FileSize = int64(wSize)

	_, err = f.Seek(0, 0)
	if err != nil {
		uc.log.Error(err)
		return nil, buguV1.ErrorInternalServerError("Internal Server Error")
	}

	metadata.FileSha1, err = pkg.FileSha1(f)
	if err != nil {
		uc.log.Error(err)
		return nil, buguV1.ErrorInternalServerError("Internal Server Error")
	}

	return uc.fr.CreateFileMetadata(ctx, metadata)
}

func (uc *ArtifactUsecase) getFile(ctx context.Context, fileId string) (*os.File, func(), error) {
	u, err := uuid.Parse(fileId)
	if err != nil {
		return nil, nil, buguV1.ErrorUuidParseFailed("parse fileId failed, err: %v", fileId)
	}

	dto, err := uc.fr.GetFileMetadata(ctx, u)
	if err != nil {
		return nil, nil, err
	}

	ok, err := pkg.PathExists(dto.FileAddr)
	if err != nil {
		return nil, nil, buguV1.ErrorUnknownError("check file failed, err: %v", err)
	}
	if !ok {
		return nil, nil, buguV1.ErrorCreateConflict("file not exist")
	}

	f, err := os.OpenFile(dto.FileAddr, os.O_RDONLY, 0)
	if err != nil {
		return nil, nil, err
	}
	return f, func() {
		err := f.Close()
		if err != nil {
			uc.log.Error(err)
		}
	}, nil
}
