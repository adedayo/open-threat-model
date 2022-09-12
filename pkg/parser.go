package otm

import (
	"io"

	"gopkg.in/yaml.v3"
)

func Parse(in io.Reader) (model OpenThreatModel, err error) {
	err = yaml.NewDecoder(in).Decode(&model)
	if err == nil {
		if valid, err := model.Validate(); !valid {
			return model, err
		}
	}
	return
}
