package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"ols/lsp"
	"strconv"
)

func EncodeMessage(msg any) string {
	content, err := json.Marshal(msg)

	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

func DecodeMessage(msg []byte) (string, []byte, error) {
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})

	if !found {
		return "", nil, errors.New("error finding header-content seperator")
	}

	contentLength, err := strconv.Atoi(string(header[len("Content-Length: "):]))

	if err != nil {
		return "", nil, err
	}

	var requestMessage lsp.RequestMessage

	if err := json.Unmarshal(content[:contentLength], &requestMessage); err != nil {
		return "", nil, err
	}

	return requestMessage.Method, content[:contentLength], nil
}

func Split(data []byte, _ bool) (advance int, token []byte, err error) {
	header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})

	if !found {
		return 0, nil, nil
	}

	contentLength, err := strconv.Atoi(string(header[len("Content-Length: "):]))

	if err != nil {
		return 0, nil, err
	}

	if len(content) < contentLength {
		// haven't read enough bytes yet
		return 0, nil, nil
	}

	totalLength := len(header) + 4 + contentLength

	return totalLength, data[:totalLength], nil
}
