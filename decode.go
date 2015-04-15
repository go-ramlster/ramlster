package ramlster

import (
	"strings"

	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v2"
)

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
	raml.Resources, err = parseResources(ramlMap, nil)
	spew.Dump(raml)
	return
}

func parseResources(ramlMap map[interface{}]interface{}, parentResouce *Resource) (resources []Resource, err error) {
	for key, value := range ramlMap {
		if k, ok := key.(string); ok {
			if isResource(k) {
				resource := Resource{}
				resource.RelativeUri = k
				if v, ok := value.(map[interface{}]interface{}); ok {
					resource.Resources, err = parseResources(v, &resource)
				}
				resources = append(resources, resource)
			} else if parentResouce != nil {
				switch k {
				case "displayName":
					parentResouce.DisplayName = value.(string)
				case "description":
					parentResouce.Description = value.(string)
				case "get", "post", "put", "delete":
					method := parseMethod(k, value)
					parentResouce.Methods = append(parentResouce.Methods, method)
				}
			}
		}
	}
	return
}

func parseMethod(name string, methodData interface{}) (method Method) {
	method.HttpMethod = HttpMethod(name)
	if methodMap, ok := methodData.(map[interface{}]interface{}); ok {
		for key, value := range methodMap {
			if k, ok := key.(string); ok {
				switch k {
				case "description":
					method.Description = value.(string)
				}
			}
		}

	}
	return
}
