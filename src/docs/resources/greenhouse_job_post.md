---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "greenhouse_job_post Resource - src"
subcategory: ""
description: |-
  
---

# greenhouse_job_post (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **active** (Boolean)
- **content** (String)
- **demographic_question_set_id** (Number)
- **external** (Boolean)
- **id** (String) The ID of this resource.
- **internal** (Boolean)
- **internal_content** (String)
- **job_id** (Number)
- **live** (Boolean)
- **location** (String)
- **questions** (Block List) (see [below for nested schema](#nestedblock--questions))

### Read-Only

- **created_at** (String)
- **first_published_at** (String)
- **updated_at** (String)

<a id="nestedblock--questions"></a>
### Nested Schema for `questions`

Required:

- **name** (String)

Read-Only:

- **active** (Boolean)
- **answer_type** (String)
- **demographic_question_set_id** (Number)
- **required** (Boolean)
- **translations** (List of Object) (see [below for nested schema](#nestedatt--questions--translations))

<a id="nestedatt--questions--translations"></a>
### Nested Schema for `questions.translations`

Read-Only:

- **language** (String)
- **name** (String)

