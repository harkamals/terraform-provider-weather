package main

import "github.com/hashicorp/terraform/helper/schema"

func resourceData() *schema.Resource {
	return &schema.Resource{
		Read: resourceDataRead,

		Schema: map[string]*schema.Schema{
			"foo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceDataRead(d *schema.ResourceData, m interface{}) error {
	d.SetId("foo")
	d.Set("foo", "bah")

	return nil
}
