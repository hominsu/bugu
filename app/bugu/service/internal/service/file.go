package service

import (
	"mime/multipart"
	nethttp "net/http"
	"path/filepath"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func (s *BuguFileService) UploadFile(ctx http.Context) error {
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

	dto, err := s.fu.SaveFile(ctx, file, filepath.Dir(s.dc.File.Path))
	if err != nil {
		return err
	}

	return ctx.JSON(nethttp.StatusOK, dto)
}

func (s *BuguFileService) DownloadFile(ctx http.Context) error {
	req := ctx.Request()
	// userid := req.FormValue("userid")

	fileID := req.FormValue("id")
	if fileID == "" {
		return errors.BadRequest("FILE_ID_EMPTY", "file id params empty")
	}

	f, cleanup, err := s.fu.GetFile(ctx, fileID)
	if err != nil {
		return err
	}
	defer cleanup()

	ctx.Response().Header().Set("Content-Type", "application/octect-stream")
	ctx.Response().Header().Set("Content-Description", "attachment;filename=\""+f.Name()+"\"")
	return ctx.Stream(nethttp.StatusOK, "application/octect-stream", f)
}
