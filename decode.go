package ramlster

import (
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v2"
)

//type raml Raml

//func (r *Raml) UnmarshalYAML(unmarshal func(interface{}) error) error {
//var localRaml raml

//if err := unmarshal(&localRaml); err != nil {
//return err
//}
//return nil
//}

func Unmarshal(data []byte) (raml Raml, err error) {
	err = yaml.Unmarshal(data, &raml)
	if err != nil {
		return
	}

	ramlMap := make(map[interface{}]interface{})
	err = yaml.Unmarshal(data, &ramlMap)
	if err != nil {
		return
	}
	err = parse(&raml, ramlMap)
	return
}

func isResource(entry string) bool {
	return strings.HasPrefix(entry, "/")
}

func parse(raml *Raml, ramlMap map[interface{}]interface{}) (err error) {
	raml.Resources = parseResources(ramlMap)
	spew.Dump(raml)
	return nil
}

func parseResources(ramlMap map[interface{}]interface{}) (resources []Resource) {
	for key, value := range ramlMap {
		if uri, ok := key.(string); ok && isResource(uri) {
			resource := Resource{}
			resource.RelativeUri = uri
			if v, ok := value.(map[interface{}]interface{}); ok {
				resource.Resources = parseResources(v)
			}
			resources = append(resources, resource)
		}
	}
	return
}

func printMap(myMap map[interface{}]interface{}, ctd int) {
	var tabs string
	for i := 0; i < ctd; i++ {
		tabs += "\t"
	}
	for key, value := range myMap {

		fmt.Printf("%v%+v:\n", tabs, key)
		v, ok := value.(map[interface{}]interface{})
		if ok {
			ctd++
			printMap(v, ctd)
		} else if v, ok := value.([]interface{}); ok {
			for _, l := range v {
				if lista, ok := l.(map[interface{}]interface{}); ok {
					printMap(lista, ctd)
				}
			}
		} else {
			fmt.Printf("%v\t%v\n", tabs, value)
		}
	}
}
