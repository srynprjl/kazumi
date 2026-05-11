package main

import (
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
	// name := audio.AudioDownload("link")
	// file_name := path.Join(misc.GetCacheDir(), name) + ".mp3"
	// file := audio.AudioSpeed(file_name, 0.85)
	// file = audio.AudioReverb(file, 1.0, 0.7, 60, 0.5)
	if checkDependencies() {
		startup()
		// cli.Execute()
	}

}
