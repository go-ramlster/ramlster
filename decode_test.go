package ramlster

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

const cartRaml = "./examples/simple-cart.raml"
const errorFormat = "%s does not match!\n Expected [%v] but was [%v]"

func TestUnmarshalSimpleCart(t *testing.T) {

	const expectedBaseUri = "http://simple-service.com"
	const expectedMediaType = "application/json"

	raml, err := Unmarshal(getRamlBytes(cartRaml))
	if err != nil {
		t.Fatal(err)
	} else if raml == nil {
		t.Fatal("returned raml shouldn't be null")
	}

	if raml.BaseUri != expectedBaseUri {
		t.Errorf(errorFormat, "BaseUri", expectedBaseUri, raml.BaseUri)
	}
	if raml.MediaType != expectedMediaType {
		t.Errorf(errorFormat, "MediaType", expectedMediaType, raml.MediaType)
	}
	protocolsSize := len(raml.Protocols)
	if protocolsSize != 2 {
		t.Errorf(errorFormat, "Protocols size", 2, protocolsSize)
	}
	//spew.Dump(raml)
}

func getRamlBytes(ramlFilePath string) []byte {
	file, err := os.Open(ramlFilePath)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	source, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return source
}
