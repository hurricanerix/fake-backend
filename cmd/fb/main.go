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
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/rs/cors"

	"github.com/hurricanerix/fake-backend/handler"
)

func main() {
	addr := flag.String("addr", ":8080", "Address to listen on.")
	defaultFileName := flag.String("default-filename", "default", "File to return if a directory is requested.")
	dataPath := flag.String("data", ".", "Directory to serve files from.")

	flag.Parse()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodHead, http.MethodGet, http.MethodPost},
		AllowCredentials: true,
	})

	dataAbsolutePath, err := filepath.Abs(*dataPath)
	if err != nil {
		panic(err)
	}

	s := &http.Server{
		Addr: *addr,
		Handler: c.Handler(handler.Fake{
			DataDir:         dataAbsolutePath,
			DefaultFileName: *defaultFileName,
		}),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("Starting server at address %s\n", *addr)
	fmt.Printf("Data Path: %s\n", dataAbsolutePath)
	fmt.Printf("Default Filename: %s\n", *defaultFileName)
	log.Fatal(s.ListenAndServe())
}
