package main

import (
	"fmt"
	"os"

	"github.com/bjornpagen/tiktok-video-processor/pkg/downloader"
	"github.com/bjornpagen/tiktok-video-processor/pkg/tiktokapi"
)

func main() {
	apiKey := os.Getenv("TIKTOK_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: TIKTOK_API_KEY environment variable not set")
		return
	}
	tiktokAPI := tiktokapi.NewTikTokAPI(apiKey)

	tiktokURL := "https://vm.tiktok.com/ZSeQ2wk3q/"
	videoURL, err := tiktokAPI.GetVideoURL(tiktokURL)
	if err != nil {
		fmt.Println("Error getting video URL:", err)
		return
	}

	videoDownloader := &downloader.VideoDownloader{}
	filename := "video.mp4"
	err = videoDownloader.DownloadVideo(videoURL, filename)
	if err != nil {
		fmt.Println("Error downloading video:", err)
		return
	}

	// TODO: Implement video processing with VideoProcessor component
	// videoprocessor.ProcessVideo(filename, "overlay_image.png", "output.mp4")

	fmt.Println("Video downloaded successfully:", filename)
}
