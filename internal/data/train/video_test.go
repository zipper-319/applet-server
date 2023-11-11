package train

import (
	"os"
	"testing"
)

func TestVideo(t *testing.T) {
	f, err := os.Open("7.wav")
	if err != nil {
		t.Error(err)
	}
	SaveVideo(f, "user1", "speaker1", "7.wav")

}

func TestStatus(t *testing.T) {
	GetTrainStatus("user1", "speaker1")
}
