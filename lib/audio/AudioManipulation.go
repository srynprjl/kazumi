package audio

import (
	"fmt"
	"os"
	p "path"
	"strings"

	"github.com/srynprjl/kazumi/lib/misc"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func AudioSpeed(path string, value float64) string {
	if value < 0.5 || value > 2.0 {
		println("Value must be between 0.5 and 2.0")
		misc.Log("Invalid value", "e")
	}
	outputs := strings.Split(path, ".mp3")
	output := outputs[0] + "_edited.mp3"
	err := ffmpeg.Input(path).Silent(true).Filter("aresample", ffmpeg.Args{"44100"}).Filter("atempo", ffmpeg.Args{fmt.Sprintf("%f", value)}).Output(output, ffmpeg.KwArgs{"audio_bitrate": "192k"}).OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		misc.Log("Error while adjusting speed", "e")
		panic(err)
	}
	misc.Log(fmt.Sprintf("Adjusted speed of audio by %.2f", value), "")
	return output
}

func AudioPitch(path string, value float64) string {
	sampleRate := 44100
	new_rate := int(float64(sampleRate) * value)
	speed_corr := 1 / value
	output := path
	if !strings.HasSuffix(path, "_edited.mp3") {
		outputs := strings.Split(path, ".mp3")
		output = outputs[0] + "_edited.mp3"
	}
	temp_output := "output.mp3"
	err := ffmpeg.Input(path).Silent(true).Filter("aresample", ffmpeg.Args{"44100"}).Filter("asetrate", ffmpeg.Args{fmt.Sprintf("%d", new_rate)}).Filter("atempo", ffmpeg.Args{fmt.Sprintf("%f", speed_corr)}).Output(temp_output).OverWriteOutput().Run()
	if err != nil {
		misc.Log("Error while adjusting pitch", "e")
		panic(err)
	}
	misc.Log(fmt.Sprintf("Adjusted pitch of audio by %.2f", value), "")
	os.Rename(temp_output, output)
	return output
}

func AudioReverb(path string, inGain float32, outGain float32, delay float32, decay float32) string {
	output := path
	if !strings.HasSuffix(path, "_edited.mp3") {
		outputs := strings.Split(path, ".mp3")
		output = outputs[0] + "_edited.mp3"
	}
	temp_output := p.Join(misc.GetCacheDir(), "output.mp3")
	print(temp_output)
	filterString := fmt.Sprintf("%.2f:%.2f:%.1f:%.2f", inGain, outGain, delay, decay)
	// fmt.Println(filterString)
	err := ffmpeg.Input(path).Filter("aresample", ffmpeg.Args{"44100"}).Filter("aecho", ffmpeg.Args{filterString}).Output(temp_output, ffmpeg.KwArgs{
		"c:a": "libmp3lame",
		"b:a": "192k",
		"vn":  "",
	}).OverWriteOutput().Silent(true).ErrorToStdOut().GlobalArgs("-hide_banner", "-loglevel", "error").Run()
	if err != nil {
		misc.Log("Error while adjusting reverb", "e")
		panic(err.Error())
	}
	misc.Log(fmt.Sprintf("Adjusted reverb of audio by inGain = %.2f, outGain = %.2f, delay = %.1f, decay = %.2f", inGain, outGain, delay, decay), "")
	os.Rename(temp_output, output)
	return output
}
