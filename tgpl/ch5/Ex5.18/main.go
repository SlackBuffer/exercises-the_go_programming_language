package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}

	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	fmt.Println(local)
	if local == "/" {
		local = "index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}

	n, err = io.Copy(f, resp.Body)

	defer func() {
		// prefer error from Copy, if any
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()

	return local, n, err
}

func main() {
	fetch("https://github.com/")
}
