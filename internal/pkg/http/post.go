package http

import (
	"bytes"
	"io"
	"net/http"
)

func Post(addr string, req []byte) ([]byte, error) {
	resp, err := http.Post(addr, "application/json", bytes.NewReader(req))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}
