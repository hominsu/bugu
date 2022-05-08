package data

import (
	"bugu/app/bugu/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
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
