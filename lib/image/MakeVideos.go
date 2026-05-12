package image

import (
	"fmt"
	"os"
	"path"

	"github.com/srynprjl/kazumi/lib/misc"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func MakeVideos(imagePath string, audioPath string) string {
	output := path.Join(misc.GetCacheDir(), "output.mp4")
	bg := ffmpeg.Input("color=c=black:s=1920x1080:d=3600", ffmpeg.KwArgs{"f": "lavfi"})
	overlayImg := ffmpeg.Input(imagePath, ffmpeg.KwArgs{"f": "mjpeg"}).
		Filter("scale", ffmpeg.Args{"-1", "480"})

	audioIn := ffmpeg.Input(audioPath)

	videoStream := bg.Overlay(overlayImg, "repeat", ffmpeg.KwArgs{
		"x": "(main_w-overlay_w)/2",
		"y": "(main_h-overlay_h)/2",
	})

	err := ffmpeg.Output(
		[]*ffmpeg.Stream{videoStream, audioIn},
		output,
		ffmpeg.KwArgs{
			"vcodec":   "libx264",
			"acodec":   "aac",
			"pix_fmt":  "yuv420p",
			"shortest": "",
			"movflags": "faststart",
		},
	).OverWriteOutput().ErrorToStdOut().Silent(true).Run()
	videoPath := fmt.Sprintf("%s.mp4", imagePath)
	if err != nil {
		panic(err)
	}
	os.Rename(output, videoPath)
	return videoPath
}
