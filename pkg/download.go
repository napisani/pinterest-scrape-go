package pkg

import (
	"bufio"
	"bytes"
	"crypto/sha1"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"
)

func DownloadImages(imageUrls []string, targetDir string, maxConcurrent int) {
	semaphore := make(chan struct{}, maxConcurrent)
	uniqueImages := map[string]bool{}
	wg := &sync.WaitGroup{}
	for _, imageUrl := range imageUrls {
		semaphore <- struct{}{}
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			defer func() { <-semaphore }()
			imgBytes, err := DownloadImage(url)
			if err != nil {
				fmt.Println(fmt.Sprintf("Error downloading %s: %s", url, err))
			}
			if _, exists := uniqueImages[hashBytes(imgBytes)]; !exists && err == nil {
				uniqueImages[hashBytes(imgBytes)] = true
				err = WriteImage(imgBytes, url, targetDir)
				if err != nil {
					fmt.Println(fmt.Sprintf("Error writing image: %s",  err))
				}
			}
		}(imageUrl)
	}
	wg.Wait()
}

func hashBytes(b []byte) string {
	return fmt.Sprintf("%x", sha1.Sum(b))
}

func WriteImage(imgBytes []byte, url string, targetDir string) error {
	base_name := urlToBaseName(url)
	output_path := path.Join(targetDir, base_name)

	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		os.MkdirAll(targetDir, os.ModePerm)
	} else if err != nil {
		return err
	}

	f, err := os.Create(output_path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(imgBytes)
	if err != nil {
		return err
	}
	return nil
}

func urlToBaseName(url string) string {
	return url[strings.LastIndex(url, "/")+1:]
}

func DownloadImage(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, err
	}

	base_name := urlToBaseName(url)
	ext := strings.ToLower(strings.Split(base_name, ".")[1])
	var imgBytes bytes.Buffer
	byteBuffer := bufio.NewWriter(&imgBytes)
	if ext == "jpg" || ext == "jpeg" {
		err = jpeg.Encode(byteBuffer, img, nil)
		if err != nil {
			return nil, err
		}
	} else if ext == "png" {
		err = png.Encode(byteBuffer, img)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("Unknown image type")
	}
	return imgBytes.Bytes(), nil
}
