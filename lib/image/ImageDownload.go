package image

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/srynprjl/kazumi/lib/misc"
)

func ImageDownload(url string, name string) {
	cache_file_name := misc.GetCacheDir()
	out, err := os.Create(path.Join(cache_file_name, name))
	misc.Log(fmt.Sprintf("Created file %s", name), "")
	if err != nil {
		misc.Log(fmt.Sprintf("Error creating %s", name), "e")
		panic(err)
	}

	resp, err := http.Get(url)
	if err != nil {
		misc.Log(fmt.Sprintf("Error accessing %s", url), "e")
		panic(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)

	if err != nil {
		misc.Log(fmt.Sprintf("Error while downloading file from %s", url), "e")
		panic(err)
	}
	misc.Log(fmt.Sprintf("Downloaded file from %s", url), "")
}
