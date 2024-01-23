package server

import (
	"applet-server/api/v2/applet"
	"applet-server/internal/data"
	jwtUtil "applet-server/internal/pkg/jwt"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"os"
)

const (
	defaultMaxMemory = 32 << 20 // 32 MB
)

const OperationVoiceDataOperationUploadFiles = "/api/v2/voice-data/file/upload"

type HandlerFormData interface {
	HandlerFormData(ctx context.Context, file *data.FileObject, req *applet.UploadFilesRequest) (*emptypb.Empty, error)
}

func RegisterFormDataHandler(s *http.Server, formService HandlerFormData) {
	r := s.Route("/")
	r.POST(OperationVoiceDataOperationUploadFiles, FormDataMind(formService))
}

func FormDataMind(srv HandlerFormData) func(ctx http.Context) error {

	return func(ctx http.Context) error {
		var in applet.UploadFilesRequest
		req := ctx.Request()
		if err := req.ParseMultipartForm(defaultMaxMemory); err != nil {
			log.Error(err)
		}
		log.Debug(ctx.Header())
		token := ctx.Header().Get(jwtUtil.AuthorizationKey)
		log.Debug(req.Form, token)
		tokenInfo, err := jwtUtil.ParseToken(token, "")
		if err != nil {
			return jwtUtil.ErrTokenInvalid
		}
		log.Debug("tokenInfo", tokenInfo)
		if err := ctx.BindForm(&in); err != nil {
			return err
		}
		file, fileHeader, _ := req.FormFile("file")
		var fileName string
		if fileHeader != nil {
			fileName = fileHeader.Filename
		}
		http.SetOperation(ctx, OperationVoiceDataOperationUploadFiles)

		out, err := srv.HandlerFormData(ctx, &data.FileObject{
			File:     file,
			FileName: fileName,
			Username: tokenInfo.Username,
		}, &in)
		if err != nil {
			return err
		}

		return ctx.Result(200, out)
	}
}

func uploadFile(ctx http.Context) error {
	req := ctx.Request()

	fileName := req.FormValue("speaker")
	file, fileHeader, err := req.FormFile("file")
	if err != nil {
		return err
	}
	fileType := req.FormValue("file_type")
	if err != nil {
		return err
	}
	defer file.Close()
	log.Info(fileType)

	f, err := os.OpenFile(fileHeader.Filename, os.O_WRONLY|os.O_CREATE, 0o666)
	if err != nil {
		return err
	}
	defer f.Close()
	_, _ = io.Copy(f, file)

	return ctx.String(200, "File "+fileName+" Uploaded successfully")
}
