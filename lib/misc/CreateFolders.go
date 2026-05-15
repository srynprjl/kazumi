package misc

import (
	"os"
)

func createIfNotExists(dir string) {
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(dir, 0755)
			if err != nil {
				panic(err)
			}
		}
	}
}

func CreateCacheFolder() {
	video_cache_dir := GetCacheDir()
	createIfNotExists(video_cache_dir)
}

func CreateHomeFolder() {
	home_video_dir := VideosDir()
	createIfNotExists(home_video_dir)

	audio_video_dir := AudioDir()
	createIfNotExists(audio_video_dir)
}
