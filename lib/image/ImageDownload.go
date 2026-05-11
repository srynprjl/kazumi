package image

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/srynprjl/kazumi/lib/misc"
)

func ImageDownload(url string, name string) {
	cache_file_name := misc.GetCacheDir()
	out, err := os.Create(path.Join(cache_file_name, name))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}
