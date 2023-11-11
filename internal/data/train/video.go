package train

import (
	"applet-server/internal/conf"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/idoubi/goz"
	"io"
	"mime/multipart"
	"net/http"
)

const url = "http://172.16.32.2:32287/mock/131"

type Train struct {
	Addr string
	*log.Helper
}

func NewTrain(config *conf.Data, logger log.Logger) *Train {
	return &Train{
		Addr:   config.Train.Addr,
		Helper: log.NewHelper(logger),
	}
}

type UploadResponse struct {
	Code    int
	Message string
	TaskId  string
	WaitNum int
}

func (t *Train) SaveVideo(videoContent io.Reader, tenantCode string, speakerId string, fileName string) (*UploadResponse, error) {
	path := "/voice/upload"

	bodyBuf := &bytes.Buffer{}
	writer := multipart.NewWriter(bodyBuf)
	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return nil, err
	}
	io.Copy(part, videoContent)

	writer.WriteField("tenant_code", tenantCode)
	writer.WriteField("speaker_id", speakerId)
	writer.WriteField("flag", "start")

	contentType := writer.FormDataContentType()
	writer.Close()
	resp, err := http.Post(fmt.Sprintf("%s%s", t.Addr, path), contentType, bodyBuf)
	if err != nil {
		return nil, err
	}
	var result UploadResponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	json.Unmarshal(body, &result)
	log.Infof("result: %#v", result)
	return &result, nil
}

type GetTrainStatusResponse struct {
	Code     int
	Message  string
	Progress int
	Status   int
}

func (t *Train) GetTrainStatus(tenantCode string, speakerId string) error {
	path := "/voice/getTrainStatus"

	resp, err := goz.Get(fmt.Sprintf("%s%s", t.Addr, path), goz.Options{
		Query: fmt.Sprintf("tenant_code=%s&speaker_id=%s", tenantCode, speakerId),
	})
	if err != nil {
		return err
	}
	body, err := resp.GetBody()
	if err != nil {
		return err
	}
	var result GetTrainStatusResponse
	json.Unmarshal(body, &result)
	log.Infof("result: %#v", result)
	return nil
}
