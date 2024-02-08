package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/napisani/pintrest-scrape-go/pkg"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	searchQuery := flag.String("search", "iphone wallpaper pinterest", "The search query to use")
	maxImages := flag.Int("max", 10, "The maximum number of images to download")
	maxConcurrent := flag.Int("concurrent", 5, "The maximum number of images to download concurrently")
	targetDir := flag.String("dir", "images", "The directory to save the images to")
	help := flag.Bool("help", false, "Show help")
	flag.Parse()

	if *help || flag.NFlag() == 0 {
		usage()
		return
	}

	args := pkg.DownloadArgs{
		SearchQuery:   *searchQuery,
		MaxImages:     *maxImages,
		MaxConcurrent: *maxConcurrent,
		TargetDir:     *targetDir,
	}

	pinterestUrls := pkg.GetPinterestUrls(args.SearchQuery)
	imageUrls := pkg.GetImageUrls(pinterestUrls, args.MaxImages)
  pkg.DownloadImages(imageUrls, args.TargetDir, args.MaxConcurrent)
  fmt.Println("Downloaded", len(imageUrls), "images")
}
