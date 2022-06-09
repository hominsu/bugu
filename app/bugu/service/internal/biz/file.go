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
	ID       uuid.UUID  `json:"fileId,omitempty"`
	FileSha1 string     `json:"fileSha1,omitempty"`
	FileSize int64      `json:"fileSize,omitempty"`
	FileAddr string     `json:"fileAddr,omitempty"`
	Type     *file.Type `json:"type,omitempty"`
}

var (
	TypeName = map[int32]file.Type{
		0:  file.TypeAdposhel,
		1:  file.TypeAgent,
		2:  file.TypeAllaple,
		3:  file.TypeAmonetize,
		4:  file.TypeAndrom,
		5:  file.TypeAutorun,
		6:  file.TypeBrowseFox,
		7:  file.TypeDinwod,
		8:  file.TypeElex,
		9:  file.TypeExpiro,
		10: file.TypeFasong,
		11: file.TypeHackKMS,
		12: file.TypeHlux,
		13: file.TypeInjector,
		14: file.TypeInstallCore,
		15: file.TypeMultiPlug,
		16: file.TypeNeoreklami,
		17: file.TypeNeshta,
		18: file.TypeOther,
		19: file.TypeRegrun,
		20: file.TypeSality,
		21: file.TypeSnarasite,
		22: file.TypeStantinko,
		23: file.TypeVBA,
		24: file.TypeVBKrypt,
		25: file.TypeVilsel,
	}

	TypeValue = map[file.Type]int32{
		file.TypeAdposhel:    0,
		file.TypeAgent:       1,
		file.TypeAllaple:     2,
		file.TypeAmonetize:   3,
		file.TypeAndrom:      4,
		file.TypeAutorun:     5,
		file.TypeBrowseFox:   6,
		file.TypeDinwod:      7,
		file.TypeElex:        8,
		file.TypeExpiro:      9,
		file.TypeFasong:      10,
		file.TypeHackKMS:     11,
		file.TypeHlux:        12,
		file.TypeInjector:    13,
		file.TypeInstallCore: 14,
		file.TypeMultiPlug:   15,
		file.TypeNeoreklami:  16,
		file.TypeNeshta:      17,
		file.TypeOther:       18,
		file.TypeRegrun:      19,
		file.TypeSality:      20,
		file.TypeSnarasite:   21,
		file.TypeStantinko:   22,
		file.TypeVBA:         23,
		file.TypeVBKrypt:     24,
		file.TypeVilsel:      25,
	}
)

type FileRepo interface {
	CreateFileMetadata(ctx context.Context, userId uuid.UUID, file *File) (*File, bool, error)
	UpdateFileMetadata(ctx context.Context, file *File) (*File, error)
	AppendFileMetadataToUser(ctx context.Context, userId, fileId uuid.UUID) (*File, error)
	GetFileMetadata(ctx context.Context, userId, fileId uuid.UUID) (*File, error)
	GetFileMetadataByUserId(ctx context.Context, userId uuid.UUID) ([]*File, error)
	DeleteFileMetadata(ctx context.Context, userId, fileId uuid.UUID) error
	QueryByFileSha1(ctx context.Context, sha1 string) (*File, error)
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

func (uc *FileUsecase) SaveFile(ctx context.Context, userId string, metaFile multipart.File, dir string) (*File, error) {
	userid, err := uuid.Parse(userId)
	if err != nil {
		return nil, buguV1.ErrorUuidParseFailed("parse userId failed, err: %v", userId)
	}

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

	sha1, size, err := pkg.IOSha1(metaFile)
	if err != nil {
		uc.log.Error(err)
		return nil, buguV1.ErrorInternalServerError("Internal Server Error")
	}

	_, err = metaFile.Seek(0, 0)
	if err != nil {
		uc.log.Error(err)
		return nil, buguV1.ErrorInternalServerError("Internal Server Error")
	}

	dto, ok, err := uc.repo.CreateFileMetadata(ctx, userid, &File{
		ID:       u,
		FileSha1: sha1,
		FileSize: size,
		FileAddr: dir,
	})
	if err != nil {
		return nil, err
	}

	if ok {
		return dto, nil
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

	_, err = io.Copy(f, metaFile)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (uc *FileUsecase) GetFile(ctx context.Context, userId, fileId string) (*os.File, func(), error) {
	userid, err := uuid.Parse(userId)
	if err != nil {
		return nil, nil, buguV1.ErrorUuidParseFailed("parse userId failed, err: %v", userId)
	}

	fid, err := uuid.Parse(fileId)
	if err != nil {
		return nil, nil, buguV1.ErrorUuidParseFailed("parse fileId failed, err: %v", fileId)
	}

	dto, err := uc.repo.GetFileMetadata(ctx, userid, fid)
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

func (uc *FileUsecase) GetFileMetadata(ctx context.Context, userId, fileId string) (*File, error) {
	userid, err := uuid.Parse(userId)
	if err != nil {
		return nil, buguV1.ErrorUuidParseFailed("parse userId failed, err: %v", userId)
	}

	fid, err := uuid.Parse(fileId)
	if err != nil {
		return nil, buguV1.ErrorUuidParseFailed("parse fileId failed, err: %v", fileId)
	}

	return uc.repo.GetFileMetadata(ctx, userid, fid)
}

func (uc *FileUsecase) GetFileMetadataByUserId(ctx context.Context, userId string) ([]*File, error) {
	userid, err := uuid.Parse(userId)
	if err != nil {
		return nil, buguV1.ErrorUuidParseFailed("parse userId failed, err: %v", userId)
	}

	return uc.repo.GetFileMetadataByUserId(ctx, userid)
}

func (uc *FileUsecase) DeleteFileMetadata(ctx context.Context, userId, fileId string) error {
	userid, err := uuid.Parse(userId)
	if err != nil {
		return buguV1.ErrorUuidParseFailed("parse userId failed, err: %v", userId)
	}

	fid, err := uuid.Parse(fileId)
	if err != nil {
		return buguV1.ErrorUuidParseFailed("parse fileId failed, err: %v", fileId)
	}

	return uc.repo.DeleteFileMetadata(ctx, userid, fid)
}
