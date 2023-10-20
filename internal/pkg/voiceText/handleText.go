package voiceText

import (
	"applet-server/api/v2/applet"
	"bufio"
	"io"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

var VoiceDataSize map[string]int

func init() {
	var (
		err     error
		content []string
	)
	VoiceDataSize = make(map[string]int, 0)
	_, fileName, _, ok := runtime.Caller(0)
	if !ok {
		panic("runtime.Caller failed")
	}
	dir := path.Dir(fileName)

	for _, v := range applet.VoiceType_name {
		content, err = ReadText(filepath.Join(dir, v))
		if err != nil {
			panic(err)
		}
		VoiceDataSize[v] = len(content)
	}
}

func ReadText(fileType string) ([]string, error) {
	f, err := os.Open(fileType + ".txt")
	if err != nil {
		return nil, err
	}
	r := bufio.NewReader(f)
	var lines []string
	for {
		lineBytes, err := r.ReadBytes('\n')
		line := strings.TrimSpace(string(lineBytes))
		if err != nil && err != io.EOF {
			return nil, err
		}
		if err == io.EOF {
			break
		}
		lines = append(lines, line)
	}
	return lines, nil
}
