package main

import (
	"fmt"
	"path"

	"github.com/srynprjl/kazumi/lib/audio"
	"github.com/srynprjl/kazumi/lib/image"
	"github.com/srynprjl/kazumi/lib/misc"
)

func checkDependencies() bool {
	a := true
	keys := [2]string{"yt-dlp", "ffmpeg"}
	for _, key := range keys {
		if bool, str := misc.CheckDependencies(key); !bool {
			println(str)
			a = false
		}
	}
	return a
}

func startup() {
	misc.CreateCacheFolder()
	misc.CreateHomeFolder()
}

func main() {
	name := audio.AudioDownload("https://music.youtube.com/watch?v=-PWJflcm2fw&si=g0J5Vbqy7tLd4zLJ")
	img := image.ImageDownload("https://4kwallpapers.com/images/wallpapers/cat-kitten-pet-domestic-animals-cute-cat-portrait-fur-baby-1280x1280-3528.jpg", name)
	audio_path := path.Join(misc.GetCacheDir(), fmt.Sprintf("%s.mp3", name))
	image.MakeVideos(img, audio_path)

	if checkDependencies() {
		startup()
		// cli.Execute()
	}

}
