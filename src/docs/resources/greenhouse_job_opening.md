---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "greenhouse_job_opening Resource - src"
subcategory: ""
description: |-
  
---

# greenhouse_job_opening (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **job_id** (Number)
- **opening_id** (String)

### Optional

- **application_id** (Number)
- **close_reason_id** (Number)
- **custom_fields** (Map of String)
- **id** (String) The ID of this resource.
- **status** (String)

### Read-Only

- **close_reason** (List of Object) (see [below for nested schema](#nestedatt--close_reason))
- **closed_at** (String)
- **opened_at** (String)

<a id="nestedatt--close_reason"></a>
### Nested Schema for `close_reason`

Read-Only:

- **id** (String)
- **name** (String)


