# pinterest-scrape-go

`pintrest-scrape-go` is a Go library that allows you to scrape and download images from Pinterest based on a search query.

## Installation

To use `pintrest-scrape-go`, you need to have Go installed and set up on your machine. Once you have Go installed, you can install the library by running the following command:

```
go get github.com/napisani/pintrest-scrape-go
```

## Usage

To use `pintrest-scrape-go` as a command-line tool, you can follow the steps below:

1. Navigate to the directory where `main.go` is located:

   ```
   cd ~/pintrest-scrape-go 
   ```

2. Build the executable:

   ```
   go build cmd/commandline/main.go 
   ```

3. Run the command-line tool with the desired options:

   ```
   ./main -search "iphone wallpaper pinterest" -max 10 -concurrent 5 -dir "images"
   ```

   The available options are:

   - `-search`: The search query to use (default: "iphone wallpaper")
   - `-max`: The maximum number of images to download (default: 10)
   - `-concurrent`: The maximum number of images to download concurrently (default: 5)
   - `-dir`: The directory to save the images to (default: "images")
   - `-help`: Show help

## Library Usage

If you want to use `pintrest-scrape-go` as a library in your Go projects, you can import it as follows:

```go
import "github.com/napisani/pintrest-scrape-go/pkg"
```

Here's an example of how to use the library to scrape and download images:

```go
args := pkg.DownloadArgs{
  SearchQuery:   "iphone wallpaper",
  MaxImages:     10,
  MaxConcurrent: 5,
  TargetDir:     "images",
}

pinterestUrls := pkg.GetPinterestUrls(args.SearchQuery)
imageUrls := pkg.GetImageUrls(pinterestUrls, args.MaxImages)
pkg.DownloadImages(imageUrls, args.TargetDir, args.MaxConcurrent)
fmt.Println("Downloaded", len(imageUrls), "images")
```

This will download the specified number of images from Pinterest based on the search query and save them in the target directory.

## Contributing

Contributions are welcome! If you find any issues or have any suggestions, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.


