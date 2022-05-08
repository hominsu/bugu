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
