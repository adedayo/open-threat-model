package main

import (
	"fmt"
	"log"
	"reflect"

	otm "github.com/adedayo/open-threat-model/pkg"
	otm2 "github.com/adedayo/open-threat-model/pkg/model"
	"gopkg.in/yaml.v3"
)

func main() {
	fmt.Printf("Version: %s\n", otm.GetVersion())

	model := otm.MakeModel()

	out, err := yaml.Marshal(model)

	if err != nil {
		log.Printf("%v", err)
		return
	}

	fmt.Printf("%v\n", string(out))

	model2 := otm2.OpenThreatModel{}

	yaml.Unmarshal(out, &model2)

	out2, err := yaml.Marshal(model)

	if err != nil {
		log.Printf("%v", err)
		return
	}

	if reflect.DeepEqual(out, out2) {
		fmt.Println("slice1 and slice2 are equal.")
	} else {
		fmt.Println("slice1 and slice2 are not equal.")
	}

}
