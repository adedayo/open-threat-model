package otm

import (
	"io"

	"gopkg.in/yaml.v3"
)

func Parse(in io.Reader) (model OpenThreatModel, err error) {
	err = yaml.NewDecoder(in).Decode(&model)
	return
}
