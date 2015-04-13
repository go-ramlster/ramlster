package ramlster

type Schema map[string]string
type HttpMethod string

type HttpHeader string
type HeaderType map[HttpHeader]NamedParameters

type HttpStatusCode int
type ResponseType map[HttpStatusCode]Response

type MediaType string
type BodyType map[MediaType]Body

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
	DisplayName string `yaml:"displayName"`
	Description string
	Resources   []Resource
	Methods     []Method
}

type Method struct {
	HttpMethod      HttpMethod
	Description     string
	Headers         HeaderType
	Protocols       []string
	QueryParameters map[string]NamedParameters `yaml:"queryParameters"`
	Body            BodyType
	Responses       ResponseType
}

type Body struct {
	FormParameters map[string]NamedParameters `yaml:"formParameters"`
	Schema         string
	Example        string
}

type Response struct {
	Description string
	Headers     HeaderType
	BodyType
}

// NamedParameters are present in the following propertires:
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
