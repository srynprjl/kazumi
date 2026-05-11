package audio

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/lrstanley/go-ytdlp"
	"github.com/srynprjl/kazumi/lib/misc"
)

func AudioDownload(link string) string {
	bool, _ := misc.CheckDependencies("yt-dlp")
	if !bool {
		misc.Log("yt-dlp not installed", "e")
		panic("yt-dlp not installed")
	}
	cache_dir, _ := os.UserCacheDir()
	temp_video_dir := path.Join(cache_dir, "kazumi", "videos")
	dl := ytdlp.New().PrintJSON().NoProgress().Format("bestaudio/best").FormatSort("ext:m4a").ExtractAudio().AudioFormat("mp3").NoOverwrites().NoPlaylist().Output(path.Join(temp_video_dir, "%(title)s.%(ext)s"))
	misc.Log("Downloading audio...", "")
	out, err := dl.Run(context.Background(), link)
	if err != nil {
		misc.Log("Error while downloading audio", "e")
		panic(err)
	}
	info, err := out.GetExtractedInfo()
	misc.Log(fmt.Sprintf("Downloaded %s", *info[0].Title), "")
	return *info[0].Title

}
