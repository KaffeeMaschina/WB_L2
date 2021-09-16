package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

/*
=== Утилита wget ===

# Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
const downloadDir = "downloads"

var (
	errorCannotGetURL     = errors.New("cannot get url")
	errorCannotCreateDir  = errors.New("cannot create directory")
	errorCannotCreateFile = errors.New("cannot create file")
)

func main() {

	url := "https://ru.wikipedia.org/wiki/Go"

	download(url)
}

func download(url string) {
	resp := GetURL(url)
	defer resp.Body.Close()

	baseURL := strings.TrimSuffix(url, filepath.Ext(url))
	basePath := filepath.Join(downloadDir, filepath.Base(baseURL))

	mkdir(basePath)

	indexFile := CreateIndexFile(filepath.Join(basePath, "index.html"))
	defer indexFile.Close()

	_, err := io.Copy(indexFile, resp.Body)
	if err != nil {
		fmt.Printf("Error copying response to file: %v\n", err)
		return
	}

	fmt.Printf("Website downloaded successfully to %s\n", basePath)
}
func GetURL(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(errorCannotGetURL)
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Wrong status code: %s\n", resp.Status)
		log.Fatal(err)
		return nil
	}
	return resp
}
func mkdir(dirPath string) {
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		log.Fatal(errorCannotCreateDir)
	}
}
func CreateIndexFile(name string) *os.File {
	indexFile, err := os.Create(name)
	if err != nil {
		log.Fatal(errorCannotCreateFile)
	}
	return indexFile
}
