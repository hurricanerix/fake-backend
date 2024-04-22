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
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

type Fake struct {
	DataDir         string
	DefaultFileName string
	TargetQueryKey  string
}

func (h Fake) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dt := time.Now()

	io.ReadAll(r.Body)

	p := strings.TrimRight(r.URL.Path, "/")
	target := path.Join(h.DataDir, p)

	if isDir(target) {
		target = path.Join(target, h.DefaultFileName)
	}

	fmt.Printf("%s %s %s <- %s\n", dt.Format("Mon, 02 Jan 2006 15:04:05 MST"), r.Method, r.URL.Path, target)
	w.Header().Add("FB-Target", target)

	if !isFile(target) {
		w.Header().Add("FB-Error", "target is not a file")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	f, err := os.Open(target)
	if err != nil {
		w.Header().Add("FB-Error", fmt.Sprintf("could not open target: %s", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	data, err := readFile(f)
	if err != nil {
		w.Header().Add("FB-Error", fmt.Sprintf("could not read target: %s", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("ETag", data.ETag)
	w.Header().Add("Content-Type", data.MimeType)
	w.WriteHeader(http.StatusOK)
	w.Write(data.Payload)
}
