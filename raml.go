package ramlster

type HttpMethod string
type HttpHeader string
type HttpStatusCode int
type QueryParameter string
type Schema map[string]string

type MediaType string

type Raml struct {
	SpecVersion   string
	Title         string
	BaseUri       string `yaml:"baseUri"`
	Version       string
	MediaType     MediaType `yaml:"mediaType"`
	Protocols     []string
	Schemas       []Schema
	Resources     []Resource
	Documentation []Documentation
}

type Documentation struct {
	Title   string
	Content string
}

type Resource struct {
	RelativeUri string
	DisplayName string
	Description string
	Resources   []Resource
	Methods     []Method
}

type Method struct {
	HttpMethod      HttpMethod
	Description     string
	Headers         map[HttpHeader]NamedParameters
	Protocols       []string
	QueryParameters map[QueryParameter]NamedParameters
	Body            map[MediaType]Body
	Responses       map[HttpStatusCode]Response
}

type Body struct {
	Schema         string
	Example        string
	FormParameters map[string]NamedParameters
}

type bodyType map[MediaType]Body
type Response struct {
	Description string
	Headers     map[HttpHeader]NamedParameters
	bodyType
}

// NamedParameters are present in the following properties:
//     * URI parameters
//     * Query string parameters
//     * Form parameters
//     * Request bodies (depending on the media type)
//     * Request headers
//     * Response headers
type NamedParameters struct {
	DisplayName string
	Description string
	Type        string
	Enum        []string
	Pattern     string
	MinLength   int
	MaxLength   int
	Minimum     int
	Maximum     int
	Example     string
	Repeat      bool
	Required    bool
	Default     string
}
