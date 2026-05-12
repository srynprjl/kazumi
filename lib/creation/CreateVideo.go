package creation

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/srynprjl/kazumi/lib/audio"
	"github.com/srynprjl/kazumi/lib/image"
	"github.com/srynprjl/kazumi/lib/misc"
)

func FullProcedure(audio_url string, image_url string, opt Options) {
	if strings.Contains(audio_url, "playlist?list") {
		panic("Playlist isn't supported yet.")
	}
	println("Downloading " + audio_url)
	audio_name, audio_path := audio.AudioDownload(audio_url)
	println("Downloaded " + audio_name + "\n")
	println("Downloaded image from " + image_url)
	image_path := image.ImageDownload(image_url, audio_name)
	print("Downloaded image at " + image_path)
	println()

	if opt.Speed.Enabled && opt.Speed.Value != 0.0 {
		fmt.Printf("\nSetting the speed to %.0f%% \n", opt.Speed.Value*100)
		audio_path = audio.AudioSpeed(audio_path, opt.Speed.Value)
	}
	if opt.Pitch.Enabled && opt.Pitch.Value != 0.0 {
		percentage := opt.Pitch.Value * 100
		fmt.Printf("\nSetting the pitch to %.0f%%\n", percentage)
		audio_path = audio.AudioPitch(audio_path, opt.Pitch.Value)
	}
	if opt.Reverb.Enabled && opt.Reverb.InGain != 0 && opt.Reverb.OutGain != 0 && opt.Reverb.Delay != 0 && opt.Reverb.Decay != 0 {
		fmt.Printf("\nSetting the audio's ingain to %.2f , outgain to %.2f, delay to %.2f and decay to %.2f", opt.Reverb.InGain, opt.Reverb.OutGain, opt.Reverb.Delay, opt.Reverb.Decay)
		audio_path = audio.AudioReverb(audio_path, opt.Reverb.InGain, opt.Reverb.OutGain, opt.Reverb.Delay, opt.Reverb.Decay)
	}

	println("Creating a video. ")
	video_path := image.MakeVideos(image_path, audio_path)
	println("Video succesfully created at " + video_path)

	println("Moving video to main videos path. ")
	new_videos_path := path.Join(misc.VideosDir(), path.Base(video_path))
	println(new_videos_path)
	err := os.Rename(video_path, new_videos_path)
	if err != nil {
		panic(err)
	}
}
