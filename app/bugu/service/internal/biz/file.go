package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

// File is the model entity for the File schema.
type File struct {
	ID       uuid.UUID `json:"id,omitempty"`
	FileHash uuid.UUID `json:"file_hash,omitempty"`
	FileSize string    `json:"file_size,omitempty"`
	FileAddr string    `json:"file_addr,omitempty"`
}

type FileRepo interface{}

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
