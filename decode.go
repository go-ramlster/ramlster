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
					method.Protocols = parseStringSlice(value.([]interface{}))
				case "headers":
					method.Headers = parseHeaders(value.(map[interface{}]interface{}))
				case "queryParameters":
					method.QueryParameters = parseQueryParameters(value.(map[interface{}]interface{}))
				case "body":
					method.Body = parseBodies(value.(map[interface{}]interface{}))
				}
			}
		}
	}
	return
}

func parseBodies(bodyData map[interface{}]interface{}) map[MediaType]Body {
	b := make(map[MediaType]Body)
	for key, value := range bodyData {
		if mediaType, ok := key.(string); ok {
			if body, ok := value.(map[interface{}]interface{}); ok {
				b[MediaType(mediaType)] = parseBody(body)
			} else {
				panic("malformated body!")
			}
		} else {
			panic("malformated bodies!")
		}
	}
	return b
}

func parseBody(bodyData map[interface{}]interface{}) (body Body) {
	for key, value := range bodyData {
		if bodyField, ok := key.(string); ok {
			switch bodyField {
			case "schema":
				body.Schema = value.(string)
			case "example":
				body.Example = value.(string)
			case "formParameters":
				body.FormParameters = parseFormParameters(value.(map[interface{}]interface{}))
			}
		}
	}
	return
}

func parseFormParameters(formData map[interface{}]interface{}) map[string]NamedParameters {
	f := make(map[string]NamedParameters)
	npMap := parseNamedParametersMap(formData)
	for i, namedParameters := range npMap {
		if formField, ok := i.(string); ok {
			f[formField] = namedParameters
		}
	}
	return f
}

func parseQueryParameters(queryParams map[interface{}]interface{}) map[QueryParameter]NamedParameters {
	q := make(map[QueryParameter]NamedParameters)
	npMap := parseNamedParametersMap(queryParams)
	for i, namedParameters := range npMap {
		if queryParam, ok := i.(string); ok {
			q[QueryParameter(queryParam)] = namedParameters
		}
	}
	return q
}

func parseHeaders(headers map[interface{}]interface{}) map[HttpHeader]NamedParameters {
	h := make(map[HttpHeader]NamedParameters)
	npMap := parseNamedParametersMap(headers)
	for i, namedParameters := range npMap {
		if header, ok := i.(string); ok {
			h[HttpHeader(header)] = namedParameters
		}
	}
	return h
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
			case "minLength":
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

func parseStringSlice(stringSlice []interface{}) []string {
	var strings []string
	for _, str := range stringSlice {
		strings = append(strings, str.(string))
	}
	return strings
}

func parseNamedParametersMap(mapInterface map[interface{}]interface{}) map[interface{}]NamedParameters {
	npMap := make(map[interface{}]NamedParameters)
	for key, value := range mapInterface {
		if params, ok := value.(map[interface{}]interface{}); ok {
			npMap[key] = parseNamedParameters(params)
		}
	}
	return npMap
}
