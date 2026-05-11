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
	}
	outputs := strings.Split(path, ".mp3")
	output := outputs[0] + "_edited.mp3"
	err := ffmpeg.Input(path).Silent(true).Filter("aresample", ffmpeg.Args{"44100"}).Filter("atempo", ffmpeg.Args{fmt.Sprintf("%f", value)}).Output(output, ffmpeg.KwArgs{"audio_bitrate": "192k"}).OverWriteOutput().Run()
	if err != nil {
		panic(err)
	}
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
		panic(err)
	}
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
	filterString := fmt.Sprintf("in_gain=%.2f:out_gain=%.2f:delays=%.1f:decays=%.2f", inGain, outGain, delay, decay)
	err := ffmpeg.Input(path).Filter("aresample", ffmpeg.Args{"44100"}).Filter("aecho", ffmpeg.Args{filterString}).Output(temp_output, ffmpeg.KwArgs{
		"c:a": "libmp3lame",
		"b:a": "192k",
		"vn":  "",
	}).OverWriteOutput().Silent(true).Run()
	if err != nil {
		panic(err.Error())
	}
	os.Rename(temp_output, output)
	return output
}
