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
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	buguV1 "github.com/hominsu/bugu/api/bugu/service/v1"
	"github.com/hominsu/bugu/app/bugu/service/internal/data/ent/file"
	"github.com/hominsu/bugu/pkg"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

// File is the model entity for the File schema.
type File struct {
	ID       uuid.UUID  `json:"id,omitempty"`
	FileSha1 string     `json:"file_sha_1,omitempty"`
	FileSize int64      `json:"file_size,omitempty"`
	FileAddr string     `json:"file_addr,omitempty"`
	Type     *file.Type `json:"type,omitempty"`
}

type FileRepo interface {
	CreateFileMetadata(ctx context.Context, file *File) (*File, error)
	UpdateFileMetadata(ctx context.Context, file *File) (*File, error)
	GetFileMetadata(ctx context.Context, id uuid.UUID) (*File, error)
}

type FileUsecase struct {
	repo FileRepo

	log *log.Helper
}

func NewFileUsecase(repo FileRepo, logger log.Logger) *FileUsecase {
	return &FileUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *FileUsecase) SaveFile(ctx context.Context, metaFile multipart.File, dir string) (*File, error) {
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

	metadata.FileSize, err = io.Copy(f, metaFile)
	if err != nil {
		return nil, err
	}

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

	return uc.repo.CreateFileMetadata(ctx, metadata)
}

func (uc *FileUsecase) GetFile(ctx context.Context, fileId string) (*os.File, func(), error) {
	u, err := uuid.Parse(fileId)
	if err != nil {
		return nil, nil, buguV1.ErrorUuidParseFailed("parse fileId failed, err: %v", fileId)
	}

	dto, err := uc.repo.GetFileMetadata(ctx, u)
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
