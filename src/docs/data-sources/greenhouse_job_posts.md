---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "greenhouse_job_posts Data Source - src"
subcategory: ""
description: |-
  
---

# greenhouse_job_posts (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **active** (Boolean)
- **id** (String) The ID of this resource.
- **job_id** (Number)
- **live** (Boolean)

### Read-Only

- **posts** (List of Object) (see [below for nested schema](#nestedatt--posts))

<a id="nestedatt--posts"></a>
### Nested Schema for `posts`

Read-Only:

- **active** (Boolean)
- **content** (String)
- **created_at** (String)
- **demographic_question_set_id** (Number)
- **external** (Boolean)
- **first_published_at** (String)
- **internal** (Boolean)
- **internal_content** (String)
- **job_id** (Number)
- **live** (Boolean)
- **questions** (List of Object) (see [below for nested schema](#nestedobjatt--posts--questions))
- **updated_at** (String)

<a id="nestedobjatt--posts--questions"></a>
### Nested Schema for `posts.questions`

Read-Only:

- **active** (Boolean)
- **answer_type** (String)
- **demographic_question_set_id** (Number)
- **name** (String)
- **required** (Boolean)
- **translations** (List of Object) (see [below for nested schema](#nestedobjatt--posts--questions--translations))

<a id="nestedobjatt--posts--questions--translations"></a>
### Nested Schema for `posts.questions.translations`

Read-Only:

- **language** (String)
- **name** (String)

