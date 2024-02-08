package pkg

import (
	"encoding/json"
	"fmt"
	"path"
	"reflect"

	"github.com/gocolly/colly"
	"golang.org/x/exp/slices"
)

type JSONRecord = map[string]interface{}

func GetImageUrls(pinterestUrls []string, maxImages int) []string {
	imageUrls := []string{}
	c := colly.NewCollector()

	c.OnHTML("script#__PWS_DATA__", func(e *colly.HTMLElement) {
		scriptText := e.Text
		urls := getImageUrls(scriptText, maxImages)
		for _, url := range urls {
      if (len(imageUrls) == maxImages){
        break
      }
			imageUrls = append(imageUrls, url)
		}
	})

	for _, pinterestUrl := range pinterestUrls {
		if len(imageUrls) < maxImages {
			c.Visit(pinterestUrl)
		}
	}
	return imageUrls
}

func getImageUrls(scriptText string, maxImages int) []string {
	urls := map[string]bool{}
	var jsonData JSONRecord
	err := json.Unmarshal([]byte(scriptText), &jsonData)
	if err != nil {
		panic(err)
	}
	pins := jsonData["props"].(JSONRecord)
	pins = pins["initialReduxState"].(JSONRecord)
	pins = pins["pins"].(JSONRecord)

	buildList := func() []string {
		urlList := []string{}
		for key, _ := range urls {
			urlList = append(urlList, key)
			if len(urlList) >= maxImages {
				break
			}
		}
		end := maxImages
		if len(urlList) < end {
			end = len(urlList)
		}
		return urlList[:end]
	}

	for _, pin := range pins {
		orig := pin.(JSONRecord)["images"].(JSONRecord)["orig"]
		if reflect.TypeOf(orig) == reflect.TypeOf([]interface{}{}) {
			for _, img := range orig.([]interface{}) {
				url := img.(JSONRecord)["url"].(string)
				if slices.Index(SupportedExtensions, path.Ext(url)) >= 0 {
					urls[url] = true
				}
				if len(urls) >= maxImages {
					return buildList()
				}
			}
		} else {
			url := orig.(JSONRecord)["url"].(string)
			if slices.Index(SupportedExtensions, path.Ext(url)) >= 0 {
				urls[url] = true
			}
			if len(urls) >= maxImages {
				return buildList()
			}
		}
	}
	return buildList()
}
