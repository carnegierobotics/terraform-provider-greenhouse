---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "greenhouse_demographic_questions Data Source - src"
subcategory: ""
description: |-
  
---

# greenhouse_demographic_questions (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **questions** (List of Object) (see [below for nested schema](#nestedatt--questions))

<a id="nestedatt--questions"></a>
### Nested Schema for `questions`

Read-Only:

- **active** (Boolean)
- **answer_type** (String)
- **demographic_question_set_id** (Number)
- **name** (String)
- **required** (Boolean)
- **translations** (List of Object) (see [below for nested schema](#nestedobjatt--questions--translations))

<a id="nestedobjatt--questions--translations"></a>
### Nested Schema for `questions.translations`

Read-Only:

- **language** (String)
- **name** (String)


