package v1

import (
	"github.com/go-kratos/kratos/v2/transport/http"
)

type BuguFileHTTPServer interface {
	UploadFile(http.Context) error
	DownloadFile(http.Context) error
}

func RegisterBuguFileHTTPServer(s *http.Server, srv BuguFileHTTPServer, filter ...http.FilterFunc) {
	r := s.Route("/", filter...)
	r.POST("/v1/{userId}/files", BuguUploadFileHTTPHandler(srv))
	r.GET("/v1/{userId}/file/{fileId}", BuguDownloadFileHTTPHandler(srv))
}

func BuguUploadFileHTTPHandler(srv BuguFileHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		return srv.UploadFile(ctx)
	}
}

func BuguDownloadFileHTTPHandler(srv BuguFileHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		return srv.DownloadFile(ctx)
	}
}
