---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "greenhouse_job Resource - src"
subcategory: ""
description: |-
  
---

# greenhouse_job (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **number_of_openings** (Number)
- **template_job_id** (Number)

### Optional

- **confidential** (Boolean)
- **custom_fields** (Map of String)
- **department_id** (Number)
- **hiring_team_id** (Number)
- **id** (String) The ID of this resource.
- **job_name** (String)
- **job_post_name** (String)
- **notes** (String)
- **office_ids** (Set of Number)
- **opening_ids** (Block Set) (see [below for nested schema](#nestedblock--opening_ids))
- **openings** (Set of Number)
- **requisition_id** (String)

### Read-Only

- **closed_at** (String)
- **copied_from_id** (Number)
- **created_at** (String)
- **departments** (Set of Object) (see [below for nested schema](#nestedatt--departments))
- **hiring_team** (Map of String)
- **is_template** (Boolean)
- **offices** (Set of Object) (see [below for nested schema](#nestedatt--offices))
- **opened_at** (String)
- **status** (String)
- **updated_at** (String)

<a id="nestedblock--opening_ids"></a>
### Nested Schema for `opening_ids`

Required:

- **opening_id** (String)

Optional:

- **close_reason_id** (Number)
- **custom_fields** (Block Set) (see [below for nested schema](#nestedblock--opening_ids--custom_fields))
- **status** (String)

Read-Only:

- **application_id** (Number)
- **closed_at** (String)
- **opened_at** (String)

<a id="nestedblock--opening_ids--custom_fields"></a>
### Nested Schema for `opening_ids.custom_fields`

Required:

- **name** (String)



<a id="nestedatt--departments"></a>
### Nested Schema for `departments`

Read-Only:

- **child_ids** (Set of Number)
- **name** (String)
- **parent_id** (Number)


<a id="nestedatt--offices"></a>
### Nested Schema for `offices`

Read-Only:

- **child_ids** (Set of Number)
- **location** (String)
- **name** (String)
- **parent_id** (Number)
- **primary_contact_user_id** (Number)

