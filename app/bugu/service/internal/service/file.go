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

package service

import (
	"context"
	"mime/multipart"
	nethttp "net/http"

	"github.com/hominsu/bugu/app/bugu/service/internal/biz"

	"github.com/go-kratos/kratos/v2/metadata"
	buguV1 "github.com/hominsu/bugu/api/bugu/service/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func (s *BuguFileService) UploadFile(ctx http.Context) error {
	userId := ctx.Vars().Get("userId")
	if userId == "" {
		return errors.Unauthorized("UNAUTHORIZED", "userId is inconsistent")
	}

	if ctx.Header().Get("x-md-global-userid") != userId {
		return errors.Unauthorized("UNAUTHORIZED", "userId is inconsistent")
	}

	req := ctx.Request()
	file, handler, err := req.FormFile("file")
	if err != nil {
		return err
	}
	defer func(file multipart.File) {
		err = file.Close()
		if err != nil {
			s.log.Error(err)
		}
	}(file)

	// 4MB
	if handler.Size > 1024*1024*4 {
		return ctx.String(nethttp.StatusBadRequest, "The file size exceeds the limit")
	}

	dto, err := s.fu.SaveFile(ctx, userId, file, s.dc.File.Path)
	if err != nil {
		return err
	}

	return ctx.JSON(nethttp.StatusOK, dto)
}

func (s *BuguFileService) DownloadFile(ctx http.Context) error {
	vars := ctx.Vars()

	userId := vars.Get("userId")
	if userId == "" {
		return errors.Unauthorized("UNAUTHORIZED", "userId is inconsistent")
	}

	if ctx.Header().Get("x-md-global-userid") != userId {
		return errors.Unauthorized("UNAUTHORIZED", "userId is inconsistent")
	}

	fileId := vars.Get("fileId")
	if fileId == "" {
		return errors.BadRequest("FILE_ID_EMPTY", "file id params empty")
	}

	f, cleanup, err := s.fu.GetFile(ctx, userId, fileId)
	if err != nil {
		return err
	}
	defer cleanup()

	ctx.Response().Header().Set("Content-Type", "application/octect-stream")
	ctx.Response().Header().Set("Content-Description", "attachment;filename=\""+f.Name()+"\"")
	return ctx.Stream(nethttp.StatusOK, "application/octect-stream", f)
}

func (s *BuguService) GetFileMeta(ctx context.Context, in *buguV1.GetFileMetaRequest) (*buguV1.GetFileMetaReply, error) {
	userId := in.GetUserId()

	md, ok := metadata.FromClientContext(ctx)
	if !ok {
		return nil, buguV1.ErrorInternalServerError("Openid does not exist in context")
	}
	if md.Get("x-md-global-userid") != userId {
		return nil, errors.Unauthorized("UNAUTHORIZED", "userid is inconsistent")
	}

	fileId := in.GetFileId()

	dto, err := s.fu.GetFileMetadata(ctx, userId, fileId)
	if err != nil {
		return nil, err
	}

	return &buguV1.GetFileMetaReply{
		FileId:    dto.ID.String(),
		FileSha_1: dto.FileSha1,
		FileSize:  dto.FileSize,
		FileAddr:  dto.FileAddr,
		Type:      buguV1.Type(biz.TypeValue[*dto.Type]),
	}, nil
}

func (s *BuguService) GetFileMetaByUserId(ctx context.Context, in *buguV1.GetFileMetaByUserIdRequest) (*buguV1.GetFileMetaByUserIdReply, error) {
	userId := in.GetUserId()

	md, ok := metadata.FromClientContext(ctx)
	if !ok {
		return nil, buguV1.ErrorInternalServerError("Openid does not exist in context")
	}
	if md.Get("x-md-global-userid") != userId {
		return nil, errors.Unauthorized("UNAUTHORIZED", "userid is inconsistent")
	}

	dtos, err := s.fu.GetFileMetadataByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	var rets []*buguV1.GetFileMetaReply
	for _, dto := range dtos {
		rets = append(rets, &buguV1.GetFileMetaReply{
			FileId:    dto.ID.String(),
			FileSha_1: dto.FileSha1,
			FileSize:  dto.FileSize,
			FileAddr:  dto.FileAddr,
			Type:      buguV1.Type(biz.TypeValue[*dto.Type]),
		})
	}

	return &buguV1.GetFileMetaByUserIdReply{FileMetadata: rets}, nil
}

func (s *BuguService) DeleteFileMetadata(ctx context.Context, in *buguV1.DeleteFileMetadataRequest) (*buguV1.DeleteFileMetadataReply, error) {
	userId := in.GetUserId()

	md, ok := metadata.FromClientContext(ctx)
	if !ok {
		return nil, buguV1.ErrorInternalServerError("Openid does not exist in context")
	}
	if md.Get("x-md-global-userid") != userId {
		return nil, errors.Unauthorized("UNAUTHORIZED", "userid is inconsistent")
	}

	fileId := in.GetFileId()

	err := s.fu.DeleteFileMetadata(ctx, userId, fileId)
	if err != nil {
		return nil, err
	}

	return &buguV1.DeleteFileMetadataReply{}, nil
}
