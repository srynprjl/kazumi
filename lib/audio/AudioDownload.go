package audio

import (
	"context"
	"log"
	"os"
	"path"

	"github.com/lrstanley/go-ytdlp"
	"github.com/srynprjl/kazumi/lib/misc"
)

func AudioDownload(link string) string {
	bool, _ := misc.CheckDependencies("yt-dlp")
	if !bool {
		log.Fatal("yt-dlp not installed")
	}
	cache_dir, _ := os.UserCacheDir()
	temp_video_dir := path.Join(cache_dir, "kazumi", "videos")
	dl := ytdlp.New().PrintJSON().NoProgress().Format("bestaudio/best").FormatSort("ext:m4a").ExtractAudio().AudioFormat("mp3").NoOverwrites().NoPlaylist().Output(path.Join(temp_video_dir, "%(title)s.%(ext)s"))
	// dl.Print
	out, err := dl.Run(context.Background(), link)
	if err != nil {
		log.Fatal(err)
	}
	info, err := out.GetExtractedInfo()
	return *info[0].Title

}
