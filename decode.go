package ramlster

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

//type raml Raml

//func (r *Raml) UnmarshalYAML(unmarshal func(interface{}) error) error {
//var localRaml raml

//if err := unmarshal(&localRaml); err != nil {
//return err
//}
//r.Title = localRaml.Title
//r.Schemas = localRaml.Schemas
//return nil
//}

//func Unmarshal(data []byte, out interface{}) (err error) {
func Unmarshal(data []byte) (raml Raml, err error) {
	err = yaml.Unmarshal(data, &raml)
	//spew.Dump(raml)
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
