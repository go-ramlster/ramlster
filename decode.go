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
	raml.Resources, err = parseResources(ramlMap, nil)
	spew.Dump(raml)
	//spew.Dump(ramlMap)
	return
}

func isResource(entry string) bool {
	return strings.HasPrefix(entry, "/")
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
				case "protocols":
					//method.Protocols = parseProtocols(value.([]interface{}))
					method.Protocols = parseStringSlice(value.([]interface{}))
				case "headers":
					method.Headers = parseHeaders(value.(map[interface{}]interface{}))
				}
			}
		}
	}
	return
}

func parseStringSlice(stringSlice []interface{}) []string {
	var strings []string
	for _, str := range stringSlice {
		strings = append(strings, str.(string))
	}
	return strings
}
func parseHeaders(headersI map[interface{}]interface{}) HeaderType {
	h := make(map[HttpHeader]NamedParameters)
	for key, value := range headersI {
		if header, ok := key.(string); ok {
			if params, ok := value.(map[interface{}]interface{}); ok {
				namedParameters := parseNamedParameters(params)
				h[HttpHeader(header)] = namedParameters
			}
		}
	}
	return HeaderType(h)
}

func parseNamedParameters(params map[interface{}]interface{}) (namedParameters NamedParameters) {
	for key, value := range params {
		if paramName, ok := key.(string); ok {
			switch paramName {
			case "displayName":
				namedParameters.DisplayName = value.(string)
			case "description":
				namedParameters.Description = value.(string)
			case "type":
				namedParameters.Type = value.(string)
			case "enum":
				namedParameters.Enum = parseStringSlice(value.([]interface{}))
			case "pattern":
				namedParameters.Pattern = value.(string)
			case "MinLength":
				namedParameters.MinLength = value.(int)
			case "maxLength":
				namedParameters.MaxLength = value.(int)
			case "minimum":
				namedParameters.Minimum = value.(int)
			case "maximum":
				namedParameters.Maximum = value.(int)
			case "example":
				namedParameters.Example = value.(string)
			case "repeat":
				namedParameters.Repeat = value.(bool)
			case "required":
				namedParameters.Required = value.(bool)
			case "default":
				namedParameters.Default = value.(string)
			}
		}
	}
	return
}
