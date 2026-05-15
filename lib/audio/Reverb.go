package audio

import (
	"fmt"
	"os"
	p "path"
	"strings"

	"github.com/srynprjl/kazumi/lib/misc"
	"github.com/srynprjl/kazumi/lib/models"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func AudioReverb(path string, reverb models.Reverb) (string, error) {
	output := path
	if !strings.HasSuffix(path, "_edited.mp3") {
		outputs := strings.Split(path, ".mp3")
		output = outputs[0] + "_edited.mp3"
	}
	temp_output := p.Join(misc.GetCacheDir(), "output.mp3")
	filterString := fmt.Sprintf("%.2f:%.2f:%.1f:%.2f", reverb.InGain, reverb.OutGain, reverb.Delay, reverb.Decay)

	fmt.Printf("Setting the audio's ingain: %.2f, outgain: %.2f, delay: %.2f and decay: %.2f",
		reverb.InGain,
		reverb.OutGain,
		reverb.Delay,
		reverb.Decay)

	err := ffmpeg.Input(path).
		Filter("aecho", ffmpeg.Args{filterString}).
		Filter("aresample", ffmpeg.Args{"44100"}).
		Output(temp_output, ffmpeg.KwArgs{
			"c:a": "libmp3lame",
			"b:a": "192k",
			"vn":  "",
		}).
		OverWriteOutput().
		Silent(true).
		ErrorToStdOut().
		GlobalArgs("-hide_banner", "-loglevel", "error", "-y").
		Run()

	if err != nil {
		return "", err
	}
	
	os.Rename(temp_output, output)
	fmt.Printf("Succesfully adjusted reverb of audio by inGain: %.2f, outGain: %.2f, delay: %.1f, decay: %.2f", reverb.InGain, reverb.OutGain, reverb.Delay, reverb.Decay)
	return output, nil
}
