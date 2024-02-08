package pkg

import (
	"encoding/json"
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
		urls := getImageUrls(scriptText, maxImages-len(imageUrls))
		for _, url := range urls {
			imageUrls = append(imageUrls, url)
		}
	})

	for _, pinterestUrl := range pinterestUrls {
		c.Visit(pinterestUrl)
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
		url_list := []string{}
		for key, _ := range urls {
			if slices.Index(SupportedExtensions, path.Ext(key)) >= 0 {
				url_list = append(url_list, key)
			}
		}
		return url_list[:maxImages]
	}

	for _, pin := range pins {
		orig := pin.(JSONRecord)["images"].(JSONRecord)["orig"]
		if reflect.TypeOf(orig) == reflect.TypeOf([]interface{}{}) {
			for _, img := range orig.([]interface{}) {
				urls[img.(JSONRecord)["url"].(string)] = true
				if len(urls) >= maxImages {
					return buildList()
				}
			}
		} else {
			urls[orig.(JSONRecord)["url"].(string)] = true
			if len(urls) >= maxImages {
				return buildList()
			}
		}
	}
	return buildList()
}
