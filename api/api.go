package api

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const API = "https://catbox.moe/user/api.php"

func Upload(s string) ([]byte, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	out := make([]byte, 0)
	bout := bytes.NewBuffer(out)
	writer := multipart.NewWriter(bout)

	s, err := filepath.Abs(s)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(s)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = writer.WriteField("reqtype", "fileupload")
	if err != nil {
		return nil, err
	}

	ff, err := writer.CreateFormFile("fileToUpload", filepath.Base(s))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(ff, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	p, err := http.NewRequest("POST", API, bout)
	if err != nil {
		return nil, err
	}
	println(writer.FormDataContentType())
	p.Header.Add("Content-Type", writer.FormDataContentType())

	resp, err := client.Do(p)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	println(bout)
	return io.ReadAll(resp.Body)
}
