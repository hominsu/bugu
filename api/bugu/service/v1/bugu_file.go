package v1

import (
	nethttp "net/http"

	"github.com/go-kratos/kratos/v2/transport/http"
)

type BuguFileServer interface {
	UploadFile(http.Context) error
	DownloadFile(http.Context) error
	mustEmbedUnimplementedBuguFileServer()
}

// UnimplementedBuguFileServer must be embedded to have forward compatible implementations.
type UnimplementedBuguFileServer struct{}

func (UnimplementedBuguFileServer) UploadFile(ctx http.Context) error {
	return ctx.String(nethttp.StatusInternalServerError, "method UploadFile not implemented")
}

func (UnimplementedBuguFileServer) DownloadFile(ctx http.Context) error {
	return ctx.String(nethttp.StatusInternalServerError, "method DownloadFile not implemented")
}

func (UnimplementedBuguFileServer) mustEmbedUnimplementedBuguFileServer() {}

type UnsafeBuguFileServer interface {
	mustEmbedUnimplementedBuguFileServer()
}
