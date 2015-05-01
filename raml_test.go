package ramlster

import "testing"

func TestGetAllResources(t *testing.T) {
	raml := buildRaml()

	resources := raml.GetResources()
	if len(resources) != 5 {
		t.Errorf(errorFormat, "Total resources count", 5, len(resources))
	}
}

func buildRaml() (raml Raml) {

	priceByIdCurrencyResource := buildResource("/price/{id}/currency", nil)
	priceByIdCurrencyResources := make([]Resource, 0)
	priceByIdCurrencyResources = append(priceByIdCurrencyResources, priceByIdCurrencyResource)

	priceByIdResources := make([]Resource, 0)
	priceByIdResources =
		append(priceByIdResources, buildResource("/price/{id}", priceByIdCurrencyResources))

	priceResource := buildResource("/price", priceByIdResources)

	cartByIdResources := make([]Resource, 0)
	cartByIdResources = append(cartByIdResources, buildResource("/cart/{id}", nil))
	cartResource := buildResource("/cart", cartByIdResources)

	ramlResources := make([]Resource, 0)
	ramlResources = append(ramlResources, priceResource)
	ramlResources = append(ramlResources, cartResource)

	raml.Resources = ramlResources
	return
}

func buildResource(uri string, subresources []Resource) (resource Resource) {
	resource.RelativeUri = uri
	resource.Resources = subresources
	return
}
