package icon

import (
	"encoding/json"
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/sync/errgroup"
)

type Icon struct {
	Size *Size  `json:"size"`
	Src  string `json:"src"`
}

type alias struct {
	Size string `json:"size"`
	Src  string `json:"src"`
}

func (i *Icon) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal(&alias{
		Src:  i.Src,
		Size: fmt.Sprintf("%vx%v", i.Size.Width, i.Size.Height),
	})
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ParseIconListFromDir(dirname string) ([]*Icon, error) {
	var g = new(errgroup.Group)

	var result []*Icon

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}
	log.Println("total files: ", len(files))

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
