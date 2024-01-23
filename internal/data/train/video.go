package train

import (
	"applet-server/api/v2/applet"
	"applet-server/internal/conf"
	"applet-server/internal/pkg/log"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/idoubi/goz"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

type Train struct {
	Addr string
	*log.MyLogger
	NotifyAddr string
}

func NewTrain(config *conf.Data, logger *log.MyLogger) *Train {
	return &Train{
		Addr:       config.Train.Addr,
		MyLogger:   logger,
		NotifyAddr: config.Train.NotifyAddr,
	}
}

type UploadResponse struct {
	Code    int
	Message string
	TaskId  string
	WaitNum int
}

func (t *Train) SaveVideo(videoContent io.Reader, tenantCode string, speakerId string, fileName string, flag applet.Flag) error {
	path := "/voice/upload"
	bodyBuf := &bytes.Buffer{}
	writer := multipart.NewWriter(bodyBuf)

	if videoContent != nil && fileName != "" {
		part, err := writer.CreateFormFile("file", fileName)
		if err != nil {
			return err
		}
		io.Copy(part, videoContent)
	}

	writer.WriteField("tenant_code", tenantCode)
	writer.WriteField("speaker_id", speakerId)

	if flag == applet.Flag_start {
		writer.WriteField("flag", "start")
	} else if flag == applet.Flag_continue {
		writer.WriteField("flag", "continue")
	} else {
		writer.WriteField("flag", "end")
	}

	contentType := writer.FormDataContentType()
	writer.Close()
	now := time.Now()
	resp, err := http.Post(fmt.Sprintf("%s%s", t.Addr, path), contentType, bodyBuf)
	if err != nil {
		return err
	}
	var result UploadResponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	json.Unmarshal(body, &result)
	log.Infof("result: %#v, cost:%dms", result, time.Since(now).Milliseconds())
	if result.Code == -1 {
		return errors.New(result.Message)
	}
	return nil
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
		Query: fmt.Sprintf("tenant_code=%s&speaker_id=%s&notify_url=%s", tenantCode, speakerId, t.NotifyAddr),
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
