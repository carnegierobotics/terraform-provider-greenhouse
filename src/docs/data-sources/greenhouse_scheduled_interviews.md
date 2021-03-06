---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "greenhouse_scheduled_interviews Data Source - src"
subcategory: ""
description: |-
  
---

# greenhouse_scheduled_interviews (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **actionable** (Boolean)
- **application_id** (Number)
- **id** (String) The ID of this resource.

### Read-Only

- **interviews** (List of Object) (see [below for nested schema](#nestedatt--interviews))

<a id="nestedatt--interviews"></a>
### Nested Schema for `interviews`

Read-Only:

- **application_id** (Number)
- **end** (List of Object) (see [below for nested schema](#nestedobjatt--interviews--end))
- **external_event_id** (String)
- **interview_id** (Number)
- **interviewers** (Set of Object) (see [below for nested schema](#nestedobjatt--interviews--interviewers))
- **location** (String)
- **organizer** (List of Object) (see [below for nested schema](#nestedobjatt--interviews--organizer))
- **start** (List of Object) (see [below for nested schema](#nestedobjatt--interviews--start))
- **status** (String)
- **video_conferencing_url** (String)

<a id="nestedobjatt--interviews--end"></a>
### Nested Schema for `interviews.end`

Read-Only:

- **date** (String)
- **date_time** (String)


<a id="nestedobjatt--interviews--interviewers"></a>
### Nested Schema for `interviews.interviewers`

Read-Only:

- **email** (String)
- **employee_id** (String)
- **first_name** (String)
- **last_name** (String)
- **name** (String)
- **response_status** (String)
- **scorecard_id** (Number)
- **user_id** (Number)


<a id="nestedobjatt--interviews--organizer"></a>
### Nested Schema for `interviews.organizer`

Read-Only:

- **created_at** (String)
- **disable_user** (Boolean)
- **disabled** (Boolean)
- **emails** (List of String)
- **employee_id** (String)
- **first_name** (String)
- **last_name** (String)
- **linked_candidate_ids** (List of Number)
- **name** (String)
- **primary_email_address** (String)
- **send_email** (Boolean)
- **site_admin** (Boolean)
- **updated_at** (String)


<a id="nestedobjatt--interviews--start"></a>
### Nested Schema for `interviews.start`

Read-Only:

- **date** (String)
- **date_time** (String)


