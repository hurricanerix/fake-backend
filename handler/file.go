// MIT License

// Copyright (c) 2024 Richard Hawkins

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package handler

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"

	"github.com/gabriel-vasile/mimetype"
)

func isFile(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}

	if info.IsDir() {
		return false
	}

	return true
}

func isDir(dirPath string) bool {
	info, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return false
	}

	if !info.IsDir() {
		return false
	}

	return true
}

type FileData struct {
	Payload  []byte
	ETag     string
	MimeType string
}

func readFile(r io.Reader) (FileData, error) {
	payload, err := io.ReadAll(r)
	if err != nil {
		return FileData{}, err
	}

	data := FileData{}
	data.Payload = payload

	h := md5.New()
	io.WriteString(h, string(payload))
	hash := fmt.Sprintf("%x", h.Sum(nil))
	data.ETag = hash

	data.MimeType = "application/octet-stream"
	if m := mimetype.Detect(payload); m != nil {
		data.MimeType = m.String()
	}

	return data, nil
}
