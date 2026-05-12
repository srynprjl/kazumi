package misc

import (
	"fmt"
	"os"
)

func CleanDir(dir string) {
	err := os.RemoveAll(dir)
	if err != nil {
		panic(err)
	}
	fmt.Println("Cleaned Successfully")
}

func CleanVideoCaches() {
	cacheDir := GetCacheDir()
	CleanDir(cacheDir)
}

func CleanLogCache() {
	cacheDir := GetLogDir()
	CleanDir(cacheDir)
}
