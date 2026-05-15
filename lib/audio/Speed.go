package audio

import (
	"errors"
	"fmt"
	"strings"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func AudioSpeed(path string, value float64) (string, error) {
	if value < 0.5 || value > 2.0 {
		return "", errors.New("The values below 0.5 and the values about 2.0 isn't supported.")
	}
	outputs := strings.Split(path, ".mp3")
	output := outputs[0] + "_edited.mp3"
	fmt.Printf("Setting the speed to %.0f%% \n", value*100)
	err := ffmpeg.Input(path).
		Silent(true).
		Filter("atempo", ffmpeg.Args{fmt.Sprintf("%f", value)}).
		Output(output, ffmpeg.KwArgs{"audio_bitrate": "192k"}).
		OverWriteOutput().
		ErrorToStdOut().
		GlobalArgs("-hide_banner", "-loglevel", "error", "-y").
		Run()

	if err != nil {
		return "", err
	}
	fmt.Printf("Successfully set the speed to %.0f%% \n", value*100)
	return output, nil
}
