---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "greenhouse_applications Data Source - src"
subcategory: ""
description: |-
  
---

# greenhouse_applications (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **created_after** (String)
- **created_before** (String)
- **id** (String) The ID of this resource.
- **job_id** (Number)
- **last_activity_after** (String)
- **status** (String)

### Read-Only

- **applications** (List of Object) (see [below for nested schema](#nestedatt--applications))

<a id="nestedatt--applications"></a>
### Nested Schema for `applications`

Read-Only:

- **advance** (Boolean)
- **answers** (List of Object) (see [below for nested schema](#nestedobjatt--applications--answers))
- **applied_at** (String)
- **attachments** (List of Object) (see [below for nested schema](#nestedobjatt--applications--attachments))
- **candidate_id** (Number)
- **credited_to** (List of Object) (see [below for nested schema](#nestedobjatt--applications--credited_to))
- **current_stage** (List of Object) (see [below for nested schema](#nestedobjatt--applications--current_stage))
- **custom_fields** (Map of String)
- **hire** (Boolean)
- **initial_stage_id** (Number)
- **job_id** (Number)
- **job_ids** (List of Number)
- **job_post_id** (Number)
- **jobs** (List of Object) (see [below for nested schema](#nestedobjatt--applications--jobs))
- **keyed_custom_fields** (List of Object) (see [below for nested schema](#nestedobjatt--applications--keyed_custom_fields))
- **last_activity_at** (String)
- **location** (List of Object) (see [below for nested schema](#nestedobjatt--applications--location))
- **prospect** (Boolean)
- **prospect_detail** (List of Object) (see [below for nested schema](#nestedobjatt--applications--prospect_detail))
- **prospect_owner_id** (Number)
- **prospect_pool_id** (Number)
- **prospect_pool_stage_id** (Number)
- **prospect_stage_id** (Number)
- **prospective_department** (List of Object) (see [below for nested schema](#nestedobjatt--applications--prospective_department))
- **prospective_department_id** (Number)
- **prospective_office** (List of Object) (see [below for nested schema](#nestedobjatt--applications--prospective_office))
- **prospective_office_id** (Number)
- **referrer** (List of Object) (see [below for nested schema](#nestedobjatt--applications--referrer))
- **reject** (Boolean)
- **rejected_at** (String)
- **rejection_details** (List of Object) (see [below for nested schema](#nestedobjatt--applications--rejection_details))
- **rejection_reason** (List of Object) (see [below for nested schema](#nestedobjatt--applications--rejection_reason))
- **source** (List of Object) (see [below for nested schema](#nestedobjatt--applications--source))
- **source_id** (Number)
- **status** (String)

<a id="nestedobjatt--applications--answers"></a>
### Nested Schema for `applications.answers`

Read-Only:

- **answer** (String)
- **question** (String)


<a id="nestedobjatt--applications--attachments"></a>
### Nested Schema for `applications.attachments`

Read-Only:

- **content** (String)
- **content_type** (String)
- **filename** (String)
- **type** (String)
- **url** (String)
- **visibility** (String)


<a id="nestedobjatt--applications--credited_to"></a>
### Nested Schema for `applications.credited_to`

Read-Only:

- **employee_id** (String)
- **first_name** (String)
- **last_name** (String)
- **name** (String)


<a id="nestedobjatt--applications--current_stage"></a>
### Nested Schema for `applications.current_stage`

Read-Only:

- **name** (String)


<a id="nestedobjatt--applications--jobs"></a>
### Nested Schema for `applications.jobs`

Read-Only:

- **anywhere** (Boolean)
- **closed_at** (String)
- **confidential** (Boolean)
- **copied_from_id** (Number)
- **created_at** (String)
- **custom_fields** (Map of String)
- **department_id** (Number)
- **departments** (List of Object) (see [below for nested schema](#nestedobjatt--applications--jobs--departments))
- **hiring_team** (List of Object) (see [below for nested schema](#nestedobjatt--applications--jobs--hiring_team))
- **how_to_sell_this_job** (String)
- **is_template** (Boolean)
- **job_name** (String)
- **job_post_name** (String)
- **notes** (String)
- **number_of_openings** (Number)
- **office_ids** (List of Number)
- **offices** (List of Object) (see [below for nested schema](#nestedobjatt--applications--jobs--offices))
- **opened_at** (String)
- **opening_ids** (List of Number)
- **openings** (List of Object) (see [below for nested schema](#nestedobjatt--applications--jobs--openings))
- **requisition_id** (String)
- **status** (String)
- **team_and_responsibilities** (String)
- **template_job_id** (Number)
- **updated_at** (String)

<a id="nestedobjatt--applications--jobs--departments"></a>
### Nested Schema for `applications.jobs.departments`

Read-Only:

- **child_department_external_ids** (List of String)
- **child_ids** (List of Number)
- **external_id** (String)
- **name** (String)
- **parent_department_external_id** (String)
- **parent_id** (Number)


<a id="nestedobjatt--applications--jobs--hiring_team"></a>
### Nested Schema for `applications.jobs.hiring_team`

Read-Only:

- **members** (List of Object) (see [below for nested schema](#nestedobjatt--applications--jobs--hiring_team--members))
- **name** (String)

<a id="nestedobjatt--applications--jobs--hiring_team--members"></a>
### Nested Schema for `applications.jobs.hiring_team.name`

Read-Only:

- **employee_id** (String)
- **first_name** (String)
- **last_name** (String)
- **name** (String)
- **responsible** (Boolean)
- **responsible_for_active_candidates** (Boolean)
- **responsible_for_future_candidates** (Boolean)
- **responsible_for_inactive_candidates** (Boolean)
- **user_id** (Number)



<a id="nestedobjatt--applications--jobs--offices"></a>
### Nested Schema for `applications.jobs.offices`

Read-Only:

- **child_ids** (List of Number)
- **location** (Map of String)
- **location_name** (String)
- **name** (String)
- **parent_id** (Number)
- **primary_contact_user_id** (Number)


<a id="nestedobjatt--applications--jobs--openings"></a>
### Nested Schema for `applications.jobs.openings`

Read-Only:

- **application_id** (Number)
- **close_reason** (Map of String)
- **close_reason_id** (Number)
- **closed_at** (String)
- **custom_fields** (Map of String)
- **opened_at** (String)
- **opening_id** (String)
- **status** (String)



<a id="nestedobjatt--applications--keyed_custom_fields"></a>
### Nested Schema for `applications.keyed_custom_fields`

Read-Only:

- **name** (String)
- **type** (String)
- **value** (String)


<a id="nestedobjatt--applications--location"></a>
### Nested Schema for `applications.location`

Read-Only:

- **address** (String)
- **name** (String)


<a id="nestedobjatt--applications--prospect_detail"></a>
### Nested Schema for `applications.prospect_detail`

Read-Only:

- **prospect_owner** (List of Object) (see [below for nested schema](#nestedobjatt--applications--prospect_detail--prospect_owner))
- **prospect_pool** (List of Object) (see [below for nested schema](#nestedobjatt--applications--prospect_detail--prospect_pool))
- **prospect_stage** (List of Object) (see [below for nested schema](#nestedobjatt--applications--prospect_detail--prospect_stage))

<a id="nestedobjatt--applications--prospect_detail--prospect_owner"></a>
### Nested Schema for `applications.prospect_detail.prospect_owner`

Read-Only:

- **name** (String)


<a id="nestedobjatt--applications--prospect_detail--prospect_pool"></a>
### Nested Schema for `applications.prospect_detail.prospect_pool`

Read-Only:

- **name** (String)


<a id="nestedobjatt--applications--prospect_detail--prospect_stage"></a>
### Nested Schema for `applications.prospect_detail.prospect_stage`

Read-Only:

- **name** (String)



<a id="nestedobjatt--applications--prospective_department"></a>
### Nested Schema for `applications.prospective_department`

Read-Only:

- **child_department_external_ids** (List of String)
- **child_ids** (List of Number)
- **external_id** (String)
- **name** (String)
- **parent_department_external_id** (String)
- **parent_id** (Number)


<a id="nestedobjatt--applications--prospective_office"></a>
### Nested Schema for `applications.prospective_office`

Read-Only:

- **child_ids** (List of Number)
- **location** (Map of String)
- **location_name** (String)
- **name** (String)
- **parent_id** (Number)
- **primary_contact_user_id** (Number)


<a id="nestedobjatt--applications--referrer"></a>
### Nested Schema for `applications.referrer`

Read-Only:

- **type** (String)
- **value** (String)


<a id="nestedobjatt--applications--rejection_details"></a>
### Nested Schema for `applications.rejection_details`

Read-Only:

- **custom_fields** (Map of String)
- **keyed_custom_fields** (List of Object) (see [below for nested schema](#nestedobjatt--applications--rejection_details--keyed_custom_fields))

<a id="nestedobjatt--applications--rejection_details--keyed_custom_fields"></a>
### Nested Schema for `applications.rejection_details.keyed_custom_fields`

Read-Only:

- **name** (String)
- **type** (String)
- **value** (String)



<a id="nestedobjatt--applications--rejection_reason"></a>
### Nested Schema for `applications.rejection_reason`

Read-Only:

- **include_defaults** (Boolean)
- **name** (String)
- **per_page** (Number)
- **type** (List of Object) (see [below for nested schema](#nestedobjatt--applications--rejection_reason--type))

<a id="nestedobjatt--applications--rejection_reason--type"></a>
### Nested Schema for `applications.rejection_reason.type`

Read-Only:

- **name** (String)



<a id="nestedobjatt--applications--source"></a>
### Nested Schema for `applications.source`

Read-Only:

- **name** (String)
- **public_name** (String)
- **type** (List of Object) (see [below for nested schema](#nestedobjatt--applications--source--type))

<a id="nestedobjatt--applications--source--type"></a>
### Nested Schema for `applications.source.type`

Read-Only:

- **name** (String)

