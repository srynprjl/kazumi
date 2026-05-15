package main

import (
	"github.com/srynprjl/kazumi/cmd/cli"
	"github.com/srynprjl/kazumi/lib/misc"
)

func checkDependencies() bool {
	a := true
	keys := [2]string{"yt-dlp", "ffmpeg"}
	for _, key := range keys {
		if bool, str := misc.CheckDependencies(key); !bool {
			println(str)
			a = false
		}
	}
	return a
}

func startup() {
	misc.CreateCacheFolder()
	misc.CreateHomeFolder()
}

func main() {
	if checkDependencies() {
		startup()
		cli.Execute()
	}

}
