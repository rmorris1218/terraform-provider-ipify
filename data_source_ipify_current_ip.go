package main

import (
	"log"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/rdegges/go-ipify"
)

func dataSourceCurrentIP() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceCurrentIPRead,
		Schema: map[string]*schema.Schema{
			"address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"address_cidr": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceCurrentIPRead(d *schema.ResourceData, m interface{}) error {
	ip, err := ipify.GetIp()
	if err != nil {
		log.Printf("[ERROR] Couldn't get my IP address: %s", err)
		d.SetId("")
		return nil
	} else {
		d.SetId(time.Now().UTC().String())
		d.Set("address", ip)
		d.Set("address_cidr", fmt.Sprintf("%s/32", ip))
	}
	return nil
}
