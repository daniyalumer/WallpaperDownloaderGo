package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/schollz/progressbar/v3"
)

func main() {
	dirName := "images"
	err := checkAndCreateDir(dirName)
	if err != nil {
		fmt.Println("Operation failed:", err)
	}

	//Creating channels and wait group
	link := "https://picsum.photos/1920/1080"
	totalRequests := 50
	concurrentRequests := 10

	ch := make(chan int, concurrentRequests)

	var wg sync.WaitGroup
	wg.Add(totalRequests)
	//Progress bar issue
	bar := progressbar.Default(int64(totalRequests))
	progressCh := make(chan int)

	go func() {
		for range progressCh {
			bar.Add(1)
		}
	}()

	for i := 0; i < totalRequests; i++ {
		ch <- 1
		go func(i int) {

			defer wg.Done()
			defer func() { <-ch }()
			fileName := fmt.Sprintf("wallpaper%d.jpg", i+1)

			err = downloadWallpaper(link, filepath.Join(dirName, fileName))
			if err != nil {
				fmt.Printf("Failed to download image: %d Error: %v", i, err)
			}

			fmt.Println("Image downloaded successfully to", filepath.Join(dirName, fileName))
			//time.Sleep(time.Second)
			progressCh <- 1
		}(i)

	}

	wg.Wait()

	fmt.Println("\nAll Downloads Complete")

}

func downloadWallpaper(link string, filePath string) error {
	resp, err := http.Get(link)
	if err != nil {
		return fmt.Errorf("failed to download image: %w", err)
	}
	defer resp.Body.Close()

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save image: %w", err)
	}
	return nil
}

func checkAndCreateDir(dirName string) error {
	_, err := os.Stat(dirName)

	if os.IsNotExist(err) {
		err := os.Mkdir(dirName, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
		}
		fmt.Println("Directory created:", dirName)
	} else if err != nil {
		fmt.Println("Error checking directory:", err)
	} else {
		fmt.Println("Directory already exists:", dirName)
	}

	return nil
}
