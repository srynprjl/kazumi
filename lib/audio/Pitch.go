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

func AudioPitch(path string, value float64) (string, error) {
	output := path
	if !strings.HasSuffix(path, "_edited.mp3") {
		outputs := strings.Split(path, ".mp3")
		output = outputs[0] + "_edited.mp3"
	}
	temp_output := p.Join(misc.GetCacheDir(), "output.mp3")
	pitchStr := strconv.FormatFloat(value, 'f', 4, 64)

	fmt.Printf("\nSetting the pitch to %.2f%%\n", value*100)
	err := ffmpeg.Input(path).
		Silent(true).
		Filter("rubberband", ffmpeg.Args{fmt.Sprintf("pitch=%s", pitchStr)}).
		Filter("aresample", ffmpeg.Args{"44100"}).
		Output(temp_output).
		OverWriteOutput().
		ErrorToStdOut().
		GlobalArgs("-hide_banner", "-loglevel", "error", "-y").
		Run()

	if err != nil {
		return "", err
	}
	
	err = os.Rename(temp_output, output)
	if err != nil {
		return "", err
	}

	fmt.Printf("\nSuccessfully set the pitch to %.2f%%\n", value*100)
	return output, nil
}
