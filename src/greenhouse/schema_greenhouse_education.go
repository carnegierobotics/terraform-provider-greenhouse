package greenhouse

import (
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseEducation() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"degree": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"degree_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"discipline": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"discipline_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"end_date": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"end_month": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"end_year": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"school_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"start_date": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"start_month": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"start_year": {
			Type:     schema.TypeInt,
			Optional: true,
		},
	}
}

func inflateEducations(list []interface{}) *[]greenhouse.Education {
	newList := make([]greenhouse.Education, len(list))
	for i := range list {
		newList[i] = list[i].(greenhouse.Education)
	}
	return &newList
}
