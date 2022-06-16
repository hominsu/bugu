// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.2.2

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type BuguHTTPServer interface {
	Confusion(context.Context, *ConfusionRequest) (*ConfusionReply, error)
	DeleteArtifactMetadata(context.Context, *DeleteArtifactMetadataRequest) (*DeleteArtifactMetadataReply, error)
	DeleteFileMetadata(context.Context, *DeleteFileMetadataRequest) (*DeleteFileMetadataReply, error)
	Detect(context.Context, *DetectRequest) (*DetectReply, error)
	GetArtifactMetadata(context.Context, *GetArtifactMetadataRequest) (*GetArtifactMetadataReply, error)
	GetArtifactMetadataByFileId(context.Context, *GetArtifactMetadataByFileIdRequest) (*GetArtifactMetadataByFileIdReply, error)
	GetCurrentUser(context.Context, *GetCurrentUserRequest) (*GetCurrentUserReply, error)
	GetFileMeta(context.Context, *GetFileMetaRequest) (*GetFileMetaReply, error)
	GetFileMetaByUserId(context.Context, *GetFileMetaByUserIdRequest) (*GetFileMetaByUserIdReply, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	Packer(context.Context, *PackerRequest) (*PackerReply, error)
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserReply, error)
}

func RegisterBuguHTTPServer(s *http.Server, srv BuguHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/users", _Bugu_Register0_HTTP_Handler(srv))
	r.POST("/v1/users/login", _Bugu_Login0_HTTP_Handler(srv))
	r.GET("/v1/user/{id}", _Bugu_GetCurrentUser0_HTTP_Handler(srv))
	r.PUT("/v1/users", _Bugu_UpdateUser0_HTTP_Handler(srv))
	r.GET("/v1/user/{user_id}/file/{file_id}/metadata", _Bugu_GetFileMeta0_HTTP_Handler(srv))
	r.GET("/v1/user/{user_id}/files", _Bugu_GetFileMetaByUserId0_HTTP_Handler(srv))
	r.DELETE("/v1/user/{user_id}/file/{file_id}/metadata", _Bugu_DeleteFileMetadata0_HTTP_Handler(srv))
	r.POST("/v1/user/file/detect", _Bugu_Detect0_HTTP_Handler(srv))
	r.GET("/v1/user/{user_id}/file/{file_id}/detect", _Bugu_Detect1_HTTP_Handler(srv))
	r.POST("/v1/user/file/confusion", _Bugu_Confusion0_HTTP_Handler(srv))
	r.GET("/v1/user/{user_id}/file/{file_id}/confusion", _Bugu_Confusion1_HTTP_Handler(srv))
	r.POST("/v1/user/file/packer", _Bugu_Packer0_HTTP_Handler(srv))
	r.GET("/v1/user/{user_id}/file/{file_id}/packer", _Bugu_Packer1_HTTP_Handler(srv))
	r.GET("/v1/user/{user_id}/artifact/{artifact_id}/metadata", _Bugu_GetArtifactMetadata0_HTTP_Handler(srv))
	r.GET("/v1/user/{user_id}/artifact/file/{file_id}", _Bugu_GetArtifactMetadataByFileId0_HTTP_Handler(srv))
	r.DELETE("/v1/user/{user_id}/artifact/{artifact_id}/metadata", _Bugu_DeleteArtifactMetadata0_HTTP_Handler(srv))
}

func _Bugu_Register0_HTTP_Handler(srv BuguHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bugu.service.v1.Bugu/Register")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterReply)
		return ctx.Result(200, reply)
	}
}

func _Bugu_Login0_HTTP_Handler(srv BuguHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bugu.service.v1.Bugu/Login")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _Bugu_GetCurrentUser0_HTTP_Handler(srv BuguHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCurrentUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bugu.service.v1.Bugu/GetCurrentUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCurrentUser(ctx, req.(*GetCurrentUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCurrentUserReply)
		return ctx.Result(200, reply)
	}
}

func _Bugu_UpdateUser0_HTTP_Handler(srv BuguHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bugu.service.v1.Bugu/UpdateUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUser(ctx, req.(*UpdateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateUserReply)
		return ctx.Result(200, reply)
	}
}

func _Bugu_GetFileMeta0_HTTP_Handler(srv BuguHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetFileMetaRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bugu.service.v1.Bugu/GetFileMeta")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetFileMeta(ctx, req.(*GetFileMetaRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetFileMetaReply)
		return ctx.Result(200, reply)
	}
}

func _Bugu_GetFileMetaByUserId0_HTTP_Handler(srv BuguHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetFileMetaByUserIdRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bugu.service.v1.Bugu/GetFileMetaByUserId")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetFileMetaByUserId(ctx, req.(*GetFileMetaByUserIdRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetFileMetaByUserIdReply)
		return ctx.Result(200, reply)
	}
}

func _Bugu_DeleteFileMetadata0_HTTP_Handler(srv BuguHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteFileMetadataRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bugu.service.v1.Bugu/DeleteFileMetadata")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteFileMetadata(ctx, req.(*DeleteFileMetadataRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteFileMetadataReply)
		return ctx.Result(200, reply)
	}
}

func _Bugu_Detect0_HTTP_Handler(srv BuguHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DetectRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bugu.service.v1.Bugu/Detect")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Detect(ctx, req.(*DetectRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DetectReply)
		return ctx.Result(200, reply)
	}
}

func _Bugu_Detect1_HTTP_Handler(srv BuguHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DetectRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bugu.service.v1.Bugu/Detect")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Detect(ctx, req.(*DetectRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DetectReply)
		return ctx.Result(200, reply)
	}
}

func _Bugu_Confusion0_HTTP_Handler(srv BuguHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ConfusionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bugu.service.v1.Bugu/Confusion")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Confusion(ctx, req.(*ConfusionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ConfusionReply)
		return ctx.Result(200, reply)
	}
}

func _Bugu_Confusion1_HTTP_Handler(srv BuguHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ConfusionRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bugu.service.v1.Bugu/Confusion")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Confusion(ctx, req.(*ConfusionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ConfusionReply)
		return ctx.Result(200, reply)
	}
}

func _Bugu_Packer0_HTTP_Handler(srv BuguHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PackerRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bugu.service.v1.Bugu/Packer")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Packer(ctx, req.(*PackerRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PackerReply)
		return ctx.Result(200, reply)
	}
}

func _Bugu_Packer1_HTTP_Handler(srv BuguHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PackerRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bugu.service.v1.Bugu/Packer")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Packer(ctx, req.(*PackerRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PackerReply)
		return ctx.Result(200, reply)
	}
}

func _Bugu_GetArtifactMetadata0_HTTP_Handler(srv BuguHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetArtifactMetadataRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bugu.service.v1.Bugu/GetArtifactMetadata")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetArtifactMetadata(ctx, req.(*GetArtifactMetadataRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetArtifactMetadataReply)
		return ctx.Result(200, reply)
	}
}

func _Bugu_GetArtifactMetadataByFileId0_HTTP_Handler(srv BuguHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetArtifactMetadataByFileIdRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bugu.service.v1.Bugu/GetArtifactMetadataByFileId")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetArtifactMetadataByFileId(ctx, req.(*GetArtifactMetadataByFileIdRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetArtifactMetadataByFileIdReply)
		return ctx.Result(200, reply)
	}
}

func _Bugu_DeleteArtifactMetadata0_HTTP_Handler(srv BuguHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteArtifactMetadataRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/bugu.service.v1.Bugu/DeleteArtifactMetadata")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteArtifactMetadata(ctx, req.(*DeleteArtifactMetadataRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteArtifactMetadataReply)
		return ctx.Result(200, reply)
	}
}

type BuguHTTPClient interface {
	Confusion(ctx context.Context, req *ConfusionRequest, opts ...http.CallOption) (rsp *ConfusionReply, err error)
	DeleteArtifactMetadata(ctx context.Context, req *DeleteArtifactMetadataRequest, opts ...http.CallOption) (rsp *DeleteArtifactMetadataReply, err error)
	DeleteFileMetadata(ctx context.Context, req *DeleteFileMetadataRequest, opts ...http.CallOption) (rsp *DeleteFileMetadataReply, err error)
	Detect(ctx context.Context, req *DetectRequest, opts ...http.CallOption) (rsp *DetectReply, err error)
	GetArtifactMetadata(ctx context.Context, req *GetArtifactMetadataRequest, opts ...http.CallOption) (rsp *GetArtifactMetadataReply, err error)
	GetArtifactMetadataByFileId(ctx context.Context, req *GetArtifactMetadataByFileIdRequest, opts ...http.CallOption) (rsp *GetArtifactMetadataByFileIdReply, err error)
	GetCurrentUser(ctx context.Context, req *GetCurrentUserRequest, opts ...http.CallOption) (rsp *GetCurrentUserReply, err error)
	GetFileMeta(ctx context.Context, req *GetFileMetaRequest, opts ...http.CallOption) (rsp *GetFileMetaReply, err error)
	GetFileMetaByUserId(ctx context.Context, req *GetFileMetaByUserIdRequest, opts ...http.CallOption) (rsp *GetFileMetaByUserIdReply, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	Packer(ctx context.Context, req *PackerRequest, opts ...http.CallOption) (rsp *PackerReply, err error)
	Register(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *RegisterReply, err error)
	UpdateUser(ctx context.Context, req *UpdateUserRequest, opts ...http.CallOption) (rsp *UpdateUserReply, err error)
}

type BuguHTTPClientImpl struct {
	cc *http.Client
}

func NewBuguHTTPClient(client *http.Client) BuguHTTPClient {
	return &BuguHTTPClientImpl{client}
}

func (c *BuguHTTPClientImpl) Confusion(ctx context.Context, in *ConfusionRequest, opts ...http.CallOption) (*ConfusionReply, error) {
	var out ConfusionReply
	pattern := "/v1/user/{user_id}/file/{file_id}/confusion"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/bugu.service.v1.Bugu/Confusion"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BuguHTTPClientImpl) DeleteArtifactMetadata(ctx context.Context, in *DeleteArtifactMetadataRequest, opts ...http.CallOption) (*DeleteArtifactMetadataReply, error) {
	var out DeleteArtifactMetadataReply
	pattern := "/v1/user/{user_id}/artifact/{artifact_id}/metadata"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/bugu.service.v1.Bugu/DeleteArtifactMetadata"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BuguHTTPClientImpl) DeleteFileMetadata(ctx context.Context, in *DeleteFileMetadataRequest, opts ...http.CallOption) (*DeleteFileMetadataReply, error) {
	var out DeleteFileMetadataReply
	pattern := "/v1/user/{user_id}/file/{file_id}/metadata"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/bugu.service.v1.Bugu/DeleteFileMetadata"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BuguHTTPClientImpl) Detect(ctx context.Context, in *DetectRequest, opts ...http.CallOption) (*DetectReply, error) {
	var out DetectReply
	pattern := "/v1/user/{user_id}/file/{file_id}/detect"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/bugu.service.v1.Bugu/Detect"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BuguHTTPClientImpl) GetArtifactMetadata(ctx context.Context, in *GetArtifactMetadataRequest, opts ...http.CallOption) (*GetArtifactMetadataReply, error) {
	var out GetArtifactMetadataReply
	pattern := "/v1/user/{user_id}/artifact/{artifact_id}/metadata"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/bugu.service.v1.Bugu/GetArtifactMetadata"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BuguHTTPClientImpl) GetArtifactMetadataByFileId(ctx context.Context, in *GetArtifactMetadataByFileIdRequest, opts ...http.CallOption) (*GetArtifactMetadataByFileIdReply, error) {
	var out GetArtifactMetadataByFileIdReply
	pattern := "/v1/user/{user_id}/artifact/file/{file_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/bugu.service.v1.Bugu/GetArtifactMetadataByFileId"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BuguHTTPClientImpl) GetCurrentUser(ctx context.Context, in *GetCurrentUserRequest, opts ...http.CallOption) (*GetCurrentUserReply, error) {
	var out GetCurrentUserReply
	pattern := "/v1/user/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/bugu.service.v1.Bugu/GetCurrentUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BuguHTTPClientImpl) GetFileMeta(ctx context.Context, in *GetFileMetaRequest, opts ...http.CallOption) (*GetFileMetaReply, error) {
	var out GetFileMetaReply
	pattern := "/v1/user/{user_id}/file/{file_id}/metadata"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/bugu.service.v1.Bugu/GetFileMeta"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BuguHTTPClientImpl) GetFileMetaByUserId(ctx context.Context, in *GetFileMetaByUserIdRequest, opts ...http.CallOption) (*GetFileMetaByUserIdReply, error) {
	var out GetFileMetaByUserIdReply
	pattern := "/v1/user/{user_id}/files"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/bugu.service.v1.Bugu/GetFileMetaByUserId"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BuguHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/v1/users/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/bugu.service.v1.Bugu/Login"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BuguHTTPClientImpl) Packer(ctx context.Context, in *PackerRequest, opts ...http.CallOption) (*PackerReply, error) {
	var out PackerReply
	pattern := "/v1/user/{user_id}/file/{file_id}/packer"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/bugu.service.v1.Bugu/Packer"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BuguHTTPClientImpl) Register(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*RegisterReply, error) {
	var out RegisterReply
	pattern := "/v1/users"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/bugu.service.v1.Bugu/Register"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BuguHTTPClientImpl) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...http.CallOption) (*UpdateUserReply, error) {
	var out UpdateUserReply
	pattern := "/v1/users"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/bugu.service.v1.Bugu/UpdateUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
