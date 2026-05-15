package misc

import (
	"fmt"
	"os"
)

func createIfNotExists(dir string) {
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(dir, 0755)
			Log(fmt.Sprintf("created folder %s", dir), "")
			if err != nil {
				panic(err)
			}
		}
	}
}

func CreateCacheFolder() {

	log_cache_dir := GetLogDir()
	createIfNotExists(log_cache_dir)
	Log("Log directory created.", "")
	//log into file that cache is created
	video_cache_dir := GetCacheDir()
	createIfNotExists(video_cache_dir)
	Log("Vido Cache directory created.", "")
	//log into file that cache is created

}

func CreateHomeFolder() {
	home_video_dir := VideosDir()
	createIfNotExists(home_video_dir)
	Log("Video directory created.", "")

	audio_video_dir := AudioDir()
	createIfNotExists(audio_video_dir)
	Log("Audio directory created.", "")
}
