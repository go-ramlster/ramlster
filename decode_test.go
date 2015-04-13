package ramlster

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestUnmarshalFile(t *testing.T) {
	//file, err := os.Open("./examples/github-api-v3.raml")
	file, err := os.Open("./examples/simple-cart.raml")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	source, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	_, err = Unmarshal(source)
	if err != nil {
		t.Fatal(err)
	}
}
