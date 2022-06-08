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

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/metadata"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	buguV1 "github.com/hominsu/bugu/api/bugu/service/v1"
)

func (s *BuguService) Confusion(ctx context.Context, in *buguV1.ConfusionRequest) (*buguV1.ConfusionReply, error) {
	userId := in.GetUserId()

	md, ok := metadata.FromClientContext(ctx)
	if !ok {
		return nil, buguV1.ErrorInternalServerError("Openid does not exist in context")
	}
	if md.Get("x-md-global-userid") != userId {
		return nil, errors.Unauthorized("UNAUTHORIZED", "userid is inconsistent")
	}

	fileId := in.GetFileId()

	dto, err := s.au.Confusion(ctx, userId, fileId)
	if err != nil {
		return nil, err
	}

	return &buguV1.ConfusionReply{
		ArtifactId:       dto.ID.String(),
		FileId:           dto.FileID.String(),
		AffiliatedFileId: dto.AffiliatedFileID.String(),
		Method:           dto.Method,
	}, nil
}

func (s *BuguService) Detect(ctx context.Context, in *buguV1.DetectRequest) (*buguV1.DetectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Detect not implemented")
}

func (s *BuguService) GetArtifactMetadata(ctx context.Context, in *buguV1.GetArtifactMetadataRequest) (*buguV1.GetArtifactMetadataReply, error) {
	userId := in.GetUserId()

	md, ok := metadata.FromClientContext(ctx)
	if !ok {
		return nil, buguV1.ErrorInternalServerError("Openid does not exist in context")
	}
	if md.Get("x-md-global-userid") != userId {
		return nil, errors.Unauthorized("UNAUTHORIZED", "userid is inconsistent")
	}

	artifactId := in.GetArtifactId()

	dto, err := s.au.GetArtifactMetadata(ctx, userId, artifactId)
	if err != nil {
		return nil, err
	}

	return &buguV1.GetArtifactMetadataReply{
		ArtifactId:       dto.ID.String(),
		FileId:           dto.FileID.String(),
		AffiliatedFileId: dto.AffiliatedFileID.String(),
		Method:           dto.Method,
	}, nil
}

func (s *BuguService) GetArtifactMetadataByFileId(ctx context.Context, in *buguV1.GetArtifactMetadataByFileIdRequest) (*buguV1.GetArtifactMetadataByFileIdReply, error) {
	userId := in.GetUserId()

	md, ok := metadata.FromClientContext(ctx)
	if !ok {
		return nil, buguV1.ErrorInternalServerError("Openid does not exist in context")
	}
	if md.Get("x-md-global-userid") != userId {
		return nil, errors.Unauthorized("UNAUTHORIZED", "userid is inconsistent")
	}

	fileId := in.GetFileId()

	artifacts, err := s.au.GetArtifactMetadataByFileId(ctx, userId, fileId)
	if err != nil {
		return nil, err
	}

	var metadataReplies []*buguV1.GetArtifactMetadataReply
	for _, artifact := range artifacts {
		metadataReplies = append(metadataReplies, &buguV1.GetArtifactMetadataReply{
			ArtifactId:       artifact.ID.String(),
			FileId:           artifact.FileID.String(),
			AffiliatedFileId: artifact.AffiliatedFileID.String(),
			Method:           artifact.Method,
		})
	}

	return &buguV1.GetArtifactMetadataByFileIdReply{ArtifactMetadata: metadataReplies}, nil
}
