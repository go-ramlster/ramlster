package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`
var y = make(map[interface{}]interface{})

func unmarshalString() {
	err := yaml.Unmarshal([]byte(data), y)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("y :\n%+v\n", y)
}

func unmarshalFile() {
	file, err := os.Open("./examples/cart-service.raml")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	source, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(source, y)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("y :\n%+v\n", y)
}

func main() {
	//unmarshalString()
	unmarshalFile()
}
