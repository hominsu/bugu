package v1

import "github.com/go-kratos/kratos/v2/transport/http"

type BuguFileHTTPServer interface {
	UploadFile(http.Context) error
	DownloadFile(http.Context) error
}

func RegisterBuguFileHTTPServer(s *http.Server, srv BuguFileHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/{userid}/files", BuguUploadFileHTTPHandler(srv))
	r.GET("/v1/{userid}/file/{id}", BuguDownloadFileHTTPHandler(srv))
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
