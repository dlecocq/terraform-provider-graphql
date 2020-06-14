package graphql

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Required: true,
				Type:     schema.TypeString,
				ForceNew: true,
			},
			"headers": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: false,
				ForceNew: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"graphql_mutation": resourceGraphqlMutation(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"graphql_query": dataSourceGraphql(),
		},
		ConfigureFunc: graphqlConfigure,
	}
}

func graphqlConfigure(d *schema.ResourceData) (interface{}, error) {
	config := &GraphqlProviderConfig{
		GQLServerUrl:   d.Get("url").(string),
		RequestHeaders: d.Get("headers").(map[string]string),
	}
	return config, nil
}

type GraphqlProviderConfig struct {
	GQLServerUrl   string
	RequestHeaders map[string]string
}