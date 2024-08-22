# Wallpaper Downloader

A Go application that downloads a specified number of images from a given URL and saves them to a local directory with controlled concurrency. This project demonstrates how to handle HTTP requests, file I/O, and concurrency in Go.

## Features

- Downloads images concurrently with a limit on the number of simultaneous requests.
- Displays a progress bar showing the download progress.
- Handles directory creation and ensures the directory exists before saving images.

## Prerequisites

- Go 1.18 or later
- `github.com/schollz/progressbar/v3` package for progress visualization

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/daniyalumer/WallpaperDownloaderGo
