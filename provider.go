package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema:         map[string]*schema.Schema{},
		ConfigureFunc:  configureProvider,

		DataSourcesMap: map[string]*schema.Resource{
			"weather_cities": resourceData(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"weather_server": resourceServer(),
		},
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	// user := d.Get("user").(string)
	// token := d.Get("token").(string)

	return nil, nil
}
