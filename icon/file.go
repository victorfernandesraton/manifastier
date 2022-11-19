package icon

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/sync/errgroup"
)

type Data struct {
	Err error
}

func ParseIconListFromDir(dirname string) ([]*Icon, error) {
	var g = new(errgroup.Group)

	var result []*Icon

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}
	fmt.Println("total files: ", len(files))

	for _, imgFile := range files {

		pathForFile := filepath.Join(dirname, imgFile.Name())

		g.Go(func() error {

			if reader, err := os.Open(pathForFile); err == nil {
				defer reader.Close()
				im, _, err := image.DecodeConfig(reader)
				if err != nil {
					return err
				}

				icon := Icon{
					Src: pathForFile,
					Size: &Size{
						Width:  uint64(im.Width),
						Height: uint64(im.Height),
					},
				}

				result = append(result, &icon)
			} else {
				return err
			}
			return nil

		})
	}

	err = g.Wait()
	if err != nil {
		return nil, err
	}

	return result, nil

}
