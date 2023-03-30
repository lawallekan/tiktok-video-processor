package main

import (
	"fmt"
	"tiktok-video-processor/pkg/downloader"
	"tiktok-video-processor/pkg/tiktokapi"
	// "tiktok-video-processor/pkg/videoprocessor"
)

func main() {
	apiKey := "your-api-key-here"
	tiktokAPI := tiktokapi.NewTikTokAPI(apiKey)

	videoID := "ZSeQ2wk3q"
	videoURL, err := tiktokAPI.GetVideoURL(videoID)
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
