package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGreenhouseSources() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceGreenhouseSourcesRead,
		Schema: map[string]*schema.Schema{
      "names": {
        Type: schema.TypeList,
        Computed: true,
        Elem: &schema.Schema{
          Type: schema.TypeString,
        },
      },
      "public_names": {
        Type: schema.TypeList,
        Computed: true,
        Elem: &schema.Schema{
          Type: schema.TypeString,
        },
      },
      "query": {
        Type: schema.TypeString,
        Optional: true,
      },
    },
	}
}

func dataSourceGreenhouseSourcesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	list, err := greenhouse.GetAllSources(meta.(*greenhouse.Client), ctx)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
  names := make([]string, len(*list), len(*list))
  public_names := make([]string, len(*list), len(*list))
  for i, source := range *list {
    names[i] = source.Name
    public_names[i] = source.PublicName
  }
  d.SetId("all")
	d.Set("names", names)
  d.Set("public_names", public_names)
	return nil
}
