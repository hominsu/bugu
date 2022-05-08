package data

import (
	"bugu/app/bugu/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.ArtifactRepo = (*artifactRepo)(nil)

type artifactRepo struct {
	data *Data
	log  *log.Helper
}

// NewArtifactRepo .
func NewArtifactRepo(data *Data, logger log.Logger) biz.ArtifactRepo {
	return &artifactRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/artifact")),
	}
}
