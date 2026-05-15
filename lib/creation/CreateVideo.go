package creation

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/srynprjl/kazumi/lib/audio"
	"github.com/srynprjl/kazumi/lib/media"
	"github.com/srynprjl/kazumi/lib/misc"
	"github.com/srynprjl/kazumi/lib/models"
)

func FullProcedure(audio_url string, image_url string, opt models.Options, needAudio bool) {
	if strings.Contains(audio_url, "playlist?list") {
		panic("Playlist isn't supported yet.")
	}
	audio_name, thumbnail, audio_path := audio.AudioDownload(audio_url)
	if opt.Speed.Enabled && opt.Speed.Value != 0.0 {
		ah, err := audio.AudioSpeed(audio_path, opt.Speed.Value)
		audio_path = ah
		if err != nil {
			fmt.Printf("[ERROR] %s", err.Error())
			return
		}
	}
	if opt.Pitch.Enabled && opt.Pitch.Value != 0.0 {
		ah, err := audio.AudioPitch(audio_path, opt.Pitch.Value)
		audio_path = ah
		if err != nil {
			fmt.Printf("[ERROR] %s", err.Error())
			return
		}
	}

	if opt.Reverb.Enabled && opt.Reverb.InGain != 0 && opt.Reverb.OutGain != 0 && opt.Reverb.Delay != 0 && opt.Reverb.Decay != 0 {
		ah, err := audio.AudioReverb(audio_path, opt.Reverb)
		audio_path = ah
		if err != nil {
			fmt.Printf("[ERROR] %s", err.Error())
			return
		}
	}

	if needAudio {
		fmt.Println("Audio has been succesfully converted. Moving to audio folder")
		audios := strings.Join([]string{audio_name, "mp3"}, ".")
		new_path := path.Join(misc.AudioDir(), audios)
		os.Rename(audio_path, new_path)
		fmt.Println("Audio has been succesfully moved to " + new_path)
		return
	}

	if image_url == "" {
		image_url = thumbnail
	}

	image_path, err := media.ImageDownload(image_url, audio_name)
	if err != nil {
		fmt.Printf("[ERROR] %s", err.Error())
		return
	}

	video_path, err := media.MakeVideos(image_path, audio_path)
	if err != nil {
		fmt.Printf("[ERROR] %s", err.Error())
		return
	}
	new_videos_path := path.Join(misc.VideosDir(), path.Base(video_path))

	fmt.Println("Video has been successfully rendered. Moving it into " + new_videos_path)
	err = os.Rename(video_path, new_videos_path)
	if err != nil {
		fmt.Println("[ERROR] " + err.Error())
	}
	fmt.Println("Video has been successfully moved to " + new_videos_path)
}
