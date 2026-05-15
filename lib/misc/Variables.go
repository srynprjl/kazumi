package misc

import (
	"os"
	"path"
)

func GetCacheDir() string {
	cache_dir, err := os.UserCacheDir()
	if err != nil {
		panic(err)
	}
	dir := path.Join(cache_dir, "kazumi")
	return dir
}

func VideosDir() string {
	home_folder, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	dir := path.Join(home_folder, "Videos", "kazumi")
	return dir
}

func AudioDir() string {
	home_folder, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	dir := path.Join(home_folder, "Music", "kazumi")
	return dir
}
