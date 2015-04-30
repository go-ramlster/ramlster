package ramlster

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestUnmarshalSimpleCart(t *testing.T) {
	const expectedBaseUri = "http://simple-service.co"
	ramlBytes, err := getRamlBytes("./examples/simple-cart.raml")
	if err != nil {
		log.Fatal(err)
	}

	raml, err := Unmarshal(ramlBytes)
	if err != nil {
		t.Fatal(err)
	} else if raml == nil {
		t.Fatal("returned raml shouldn't be null")
	}
	//spew.Dump(raml)
	if raml.BaseUri != expectedBaseUri {
		t.Errorf("BaseUri does not match!\n Expected: [%s] but got: [%s]", expectedBaseUri, raml.BaseUri)
	}
	//fmt.Printf("title: %+v\n", raml.Title)
	//fmt.Printf("baseUri: %+v\n", raml.BaseUri)
	//fmt.Printf("version: %+v\n", raml.Version)
	//fmt.Printf("mediaType: %+v\n", raml.MediaType)
	//fmt.Printf("protocols: %+v\n", raml.Protocols)
	//for _, resource := range raml.Resources {
	//fmt.Printf("RelativeUri: %+v\n", resource.RelativeUri)
	//fmt.Printf("displayName: %+v\n", resource.DisplayName)
	//fmt.Printf("description: %+v\n", resource.Description)
	//}
}

func getRamlBytes(ramlFilePath string) ([]byte, error) {
	file, err := os.Open(ramlFilePath)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	source, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return source, nil
}
