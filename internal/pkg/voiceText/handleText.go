package voiceText

import (
	"applet-server/api/v2/applet"
	"bufio"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

var VoiceDataSize map[applet.VoiceType]int
var voiceDir string

func init() {
	var (
		err     error
		content []string
	)
	VoiceDataSize = make(map[applet.VoiceType]int, 0)
	_, fileName, _, ok := runtime.Caller(0)
	if !ok {
		panic("runtime.Caller failed")
	}
	voiceDir = path.Dir(fileName)

	for _, v := range applet.VoiceType_value {
		content, err = ReadText(applet.VoiceType(v))
		if err != nil {
			panic(err)
		}
		VoiceDataSize[applet.VoiceType(v)] = len(content)
	}
}

func ReadText(fileType applet.VoiceType) ([]string, error) {
	fileName := filepath.Join(voiceDir, fileType.String()+".txt")
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}
	return lines, nil
}
