---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "greenhouse_offer Data Source - src"
subcategory: ""
description: |-
  
---

# greenhouse_offer (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **application_id** (Number)
- **candidate_id** (Number)
- **custom_fields** (Map of String)
- **id** (Number) The ID of this resource.
- **job_id** (Number)
- **keyed_custom_fields** (List of Map of Object)
- **opening** (Block List, Max: 1) (see [below for nested schema](#nestedblock--opening))
- **status** (String)

### Read-Only

- **created_at** (String)
- **resolved_at** (String)
- **sent_at** (String)
- **starts_at** (String)
- **version** (Number)

<a id="nestedblock--opening"></a>
### Nested Schema for `opening`

Required:

- **opening_id** (String)

Optional:

- **application_id** (Number)
- **close_reason_id** (Number)
- **custom_fields** (Map of String)
- **status** (String)

Read-Only:

- **close_reason** (Map of String)
- **closed_at** (String)
- **opened_at** (String)

