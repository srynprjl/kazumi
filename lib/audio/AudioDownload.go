package audio

import (
	"context"
	"fmt"
	"path"
	"strings"

	"github.com/lrstanley/go-ytdlp"
	"github.com/srynprjl/kazumi/lib/misc"
)

func AudioDownload(link string) (string, string, string) {
	temp_video_dir := misc.GetCacheDir()
	dl := ytdlp.
		New().
		PrintJSON().
		NoProgress().
		Format("bestaudio/best").
		FormatSort("ext:m4a").
		ExtractAudio().
		AudioFormat("mp3").
		NoOverwrites().
		NoPlaylist().
		Output(path.Join(temp_video_dir, "%(title)s.%(ext)s"))

	fmt.Printf("Downloading audio from %s\n", link)
	out, err := dl.Run(context.Background(), link)
	if err != nil {
		fmt.Println("Error while downloading audio", err.Error())
	}
	fmt.Printf("Downloaded audio from %s\n", link)
	info, _ := out.GetExtractedInfo()
	output := strings.TrimSuffix(*info[0].Filename, "m4a") + "mp3"
	
	return *info[0].Title, *info[0].Thumbnail, output

}
