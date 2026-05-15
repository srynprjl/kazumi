package creation

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ParseJSON(jsonfile string) []JSONConfig {
	_, err := os.Stat(jsonfile)
	if os.IsNotExist(err) {
		fmt.Println("ERROR: File not found")
		os.Exit(1)
	}
	jsonFile, err := os.Open(jsonfile)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	jsonByte, err := io.ReadAll(jsonFile)
	var contents []JSONConfig
	err = json.Unmarshal(jsonByte, &contents)

	return contents
}

func DownloadUsingJSON(json []JSONConfig) {
	for _, val := range json {
		FullProcedure(val.VideoURL, val.ImageURL, Options{val.Speed, val.Pitch, val.Reverb}, val.Video)
	}
}
