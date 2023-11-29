package train

import (
	"applet-server/api/v2/applet"
	"github.com/go-kratos/kratos/v2/log"
	"os"
	"testing"
)

func TestVideo(t *testing.T) {
	train := &Train{
		Addr:   url,
		Helper: log.NewHelper(log.DefaultLogger),
	}
	fileName := "7.wav"
	f, _ := os.OpenFile(fileName, os.O_RDONLY, 0666)
	err := train.SaveVideo(f, "test", "1", fileName, applet.Flag_start)
	if err != nil {
		t.Error(err)
	}

}
