package service

import (
	buguV1 "bugu/api/bugu/service/v1"
	"bugu/app/bugu/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewBuguService)

type BuguService struct {
	buguV1.UnimplementedBuguServer

	uu  *biz.UserUsecase
	log *log.Helper
}

func NewBuguService(uu *biz.UserUsecase, logger log.Logger) *BuguService {
	return &BuguService{
		uu:  uu,
		log: log.NewHelper(log.With(logger, "module", "service/bugu")),
	}
}
