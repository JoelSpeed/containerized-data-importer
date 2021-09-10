package openapi

import (
	"github.com/emicklei/go-restful"

	"k8s.io/kube-openapi/pkg/builder"
	"k8s.io/kube-openapi/pkg/common"
	"k8s.io/kube-openapi/pkg/validation/spec"
)

// code stolen/adapted from https://github.com/kubevirt/kubevirt/blob/master/pkg/util/openapi/openapi.go

func createConfig(getDefinitions common.GetOpenAPIDefinitions) *common.Config {
	return &common.Config{
		CommonResponses: map[int]spec.Response{
			401: {
				ResponseProps: spec.ResponseProps{
					Description: "Unauthorized",
				},
			},
		},
		Info: &spec.Info{
			InfoProps: spec.InfoProps{
				Title:       "KubeVirt Containerized Data Importer API",
				Description: "Containerized Data Importer for KubeVirt.",
				Contact: &spec.ContactInfo{
					Name:  "kubevirt-dev",
					Email: "kubevirt-dev@googlegroups.com",
					URL:   "https://github.com/kubevirt/containerized-data-importer",
				},
				License: &spec.License{
					Name: "Apache 2.0",
					URL:  "https://www.apache.org/licenses/LICENSE-2.0",
				},
			},
		},
		SecurityDefinitions: &spec.SecurityDefinitions{
			"BearerToken": &spec.SecurityScheme{
				SecuritySchemeProps: spec.SecuritySchemeProps{
					Type:        "apiKey",
					Name:        "authorization",
					In:          "header",
					Description: "Bearer Token authentication",
				},
			},
		},
		GetDefinitions: getDefinitions,
	}
}

// LoadOpenAPISpec creates a swagger doc for given webservice(s)
func LoadOpenAPISpec(webServices []*restful.WebService, getDefinitions common.GetOpenAPIDefinitions) (*spec.Swagger, error) {
	config := createConfig(getDefinitions)
	return builder.BuildOpenAPISpec(webServices, config)
}
