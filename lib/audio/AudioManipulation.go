package audio

import (
	"fmt"
	"os"
	p "path"
	"strconv"
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

	err := ffmpeg.Input(path).
		Silent(true).
		Filter("atempo", ffmpeg.Args{fmt.Sprintf("%f", value)}).
		Output(output, ffmpeg.KwArgs{"audio_bitrate": "192k"}).
		OverWriteOutput().
		ErrorToStdOut().
		GlobalArgs("-hide_banner", "-loglevel", "error", "-y").
		Run()

	if err != nil {
		misc.Log("Error while adjusting speed", "e")
		panic(err)
	}
	misc.Log(fmt.Sprintf("Adjusted speed of audio by %.2f", value), "")
	return output
}

func AudioPitch(path string, value float64) string {

	output := path
	if !strings.HasSuffix(path, "_edited.mp3") {
		outputs := strings.Split(path, ".mp3")
		output = outputs[0] + "_edited.mp3"
	}
	temp_output := p.Join(misc.GetCacheDir(), "output.mp3")
	pitchStr := strconv.FormatFloat(value, 'f', 4, 64)

	err := ffmpeg.Input(path).
		Silent(true).
		Filter("rubberband", ffmpeg.Args{fmt.Sprintf("pitch=%s", pitchStr)}).
		Filter("aresample", ffmpeg.Args{"44100"}).
		Output(temp_output).
		OverWriteOutput().
		GlobalArgs("-hide_banner", "-loglevel", "error", "-y").
		Run()

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
	filterString := fmt.Sprintf("%.2f:%.2f:%.1f:%.2f", inGain, outGain, delay, decay)
	// fmt.Println(filterString)
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
		misc.Log("Error while adjusting reverb", "e")
		panic(err.Error())
	}
	misc.Log(fmt.Sprintf("Adjusted reverb of audio by inGain = %.2f, outGain = %.2f, delay = %.1f, decay = %.2f", inGain, outGain, delay, decay), "")
	os.Rename(temp_output, output)
	return output
}
