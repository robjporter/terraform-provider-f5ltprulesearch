package f5ltprulesearch

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"response": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("F5LTPRULES_RESPONSE", nil),
			},
			"rulename": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("F5LTPRULES_RULENAME", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"f5ltprules_order": resourceOrder(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}


func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	response := d.Get("response").(string)
	rulename := d.Get("rulename").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if (response != "") && (rulename != "") {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create HashiCups client",
			Detail:   "Unable to authenticate user for authenticated HashiCups client",
		})
	}

	return nil, diags
}

func resourceOrder() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOrderCreate,
		ReadContext:   resourceOrderRead,
		UpdateContext: resourceOrderUpdate,
		DeleteContext: resourceOrderDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceOrderCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	fmt.Println("resourceOrderCreate")
	return diags
}

func resourceOrderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	fmt.Println("resourceOrderRead")
	return nil
}

func resourceOrderUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	fmt.Println("resourceOrderUpdate")
	return nil
}

func resourceOrderDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	fmt.Println("resourceOrderDelete")
	return nil
}
