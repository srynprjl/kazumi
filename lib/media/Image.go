package media

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/srynprjl/kazumi/lib/misc"
)

func ImageDownload(url string, name string) (string, error) {
	cache_file_name := misc.GetCacheDir()

	fmt.Printf("Downloading image from %s\n", url)
	image_name := strings.Join([]string{name, "jpg"}, ".")
	out, err := os.Create(path.Join(cache_file_name, image_name))
	if err != nil {
		return "", err
	}
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Printf("Successfully downloaded image from %s\n", url)
	return out.Name(), nil
}
