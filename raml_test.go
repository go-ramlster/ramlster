package ramlster

import "testing"

func TestGetAllResources(t *testing.T) {
	raml := buildRaml()
	totalResources := len(raml.GetResources())

	if totalResources != 5 {
		t.Errorf(errorFormat, "Total resources count", 5, totalResources)
	}
}

func buildRaml() (raml Raml) {
	priceResource := buildResource("/price")
	priceByIdResource := buildResource("/price/{id}")
	priceByIdCurrencyResource := buildResource("/price/{id}/currency")
	addSubResource(priceByIdResource, priceByIdCurrencyResource)
	addSubResource(priceResource, priceByIdResource)

	cartResource := buildResource("/cart")
	cartByIdResource := buildResource("/cart/{id}")
	addSubResource(cartResource, cartByIdResource)

	raml.Resources = []Resource{*priceResource, *cartResource}

	return
}

func buildResource(uri string) (resource *Resource) {
	resource = new(Resource)
	resource.RelativeUri = uri
	return
}

func addSubResource(parent, child *Resource) {
	if len(parent.Resources) > 0 {
		parent.Resources = append(parent.Resources, *child)
	} else {
		resources := make([]Resource, 0)
		resources = append(resources, *child)
		parent.Resources = resources
	}
}
