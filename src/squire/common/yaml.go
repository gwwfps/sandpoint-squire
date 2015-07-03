package common

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func UnmarshalYAML(filePath string, store interface{}) {
	content, err := ioutil.ReadFile("./data/" + filePath)

	if err == nil {
		err = yaml.Unmarshal(content, store)
	}

	if err != nil {
		log.Fatalln("Cannot read data file", filePath, err)
	}

	log.Printf("Data file %s loaded.\n", filePath)
}
