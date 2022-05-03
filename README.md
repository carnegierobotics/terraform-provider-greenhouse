# terraform-provider-greenhouse

## Using the provider
Generated docs can be found at [src/docs](https://github.com/carnegierobotics/terraform-provider-greenhouse/tree/main/src/docs), but they're not super-readable in GitHub markdown. Below is a summary of how to manage some resource types.
### Provider setup
***Note:*** this is not necessary if you're using the pipeline to run Terraform.

|Attribute|Type|Required|Default|Description|
|---------|----|--------|-------|-----------|
|harvest_token|String|yes|-|Token for the Harvest API|
|jobs_token|String|yes|-|Token for the Job Board API|
|on_behalf_of|String|yes|-|User ID on whose behalf requests are made|
|harvest_url|String|no||Harvest API URL|
|jobs_url|String|no||Job Board API URL|
### Departments
|Attribute|Type|Required|Default|Description|
|---------|----|--------|-------|-----------|
|name|String|yes|-|The department's name|
|parent_id|Integer|no|-|The ID of this department's parent department|

Example:
```
resource "greenhouse_department" "mydepartment" {
  name = "My department"
  parent_id = greenhouse_department.myparentdepartment.id
}
```
### Offices

|Attribute|Type|Required|Default|Description|
|---------|----|--------|-------|-----------|
|name|String|yes|-|The office's name|
|location_name|String|no|-|A free-text field for the office's location|
|parent_id|Integer|no|-|The ID of this office's parent office|
|primary_contact_user_id|Integer|no|-|The ID of this office's primary contact user|

Example:
```
resource "greenhouse_office" "myoffice" {
  name = "My office"
  location_name = "Right here"
  parent_id = greenhouse_office.myparentoffice.id
  primary_contact_user_id = greenhouse_user.bigbossperson.id
}
```
### Users

|Attribute|Type|Required|Default|Description|
|---------|----|--------|-------|-----------|
|first_name|String|yes|-|The user's first name|
|last_name|String|yes|-|The user's last name|
|primary_email_address|String|yes|-|The user's email address|
|disable_user|Boolean|no|false|If true, disable this user|
|employee_id|String|no|-|Free-text for an internal employee ID|
|send_email|Boolean|no|false|Send this user a notification email|

Example:
```
resource "greenhouse_user" "bigbossperson" {
  first_name = "Boss"
  last_name = "Person"
  primary_email_address = "bigbossperson@mycompany.com"
  disable_user = false
  employee_id = "001"
  send_email = false
}
```
### Jobs

|Attribute|Type|Required|Default|Description|
|---------|----|--------|-------|-----------|
|number_of_openings|Integer|yes|-|Number of openings for this job|
|template_job_id|Integer|yes|-|The ID of the template job to use|
|anywhere|Boolean|no|-|This job can be done from anywhere|
|confidential|Boolean|no|-|This job is confidential|
|custom_fields|Map of string|no|-|Custom fields to attach|
|department_id|integer|no|-|The ID of the job with which this job is associated|
|hiring_team|Block set (see below)|no|-|The hiring team for this job|
|how_to_sell_this_job|String|no|-|Entry for the "How to sell this job" field|
|job_name|String|no|If not specified, it becomes "Copy of \<template job name\>"|The name for this job|
|job_post_name|String|no|-|The name to be used for the job post|
|notes|String|no|-|Notes about this job|
|office_ids|Set of Integer|no|-|A list of office ids to associate with this job|
|opening_ids|Set of Integer|no|-|A list of job opening IDs to associate with this job|
|requisition_id|String|no|-|Requisition ID to use for this job|
|team_and_responsibilities|String|no|-|Fills in the "Team and Responsibilities" field|
Hiring team block set:
|Attribute|Type|Required|Default|Description|
|---------|----|--------|-------|-----------|
|name|String|yes|-|The hiring team's name|
|members|Block set (see below)|yes|-|A list of hiring team members|
Hiring team member block set:
|Attribute|Type|Required|Default|Description|
|---------|----|--------|-------|-----------|
|user_id|Integer|yes|-|The user's ID number|
|responsible_for_active_candidates|Boolean|no|false|The user is responsible for active candidates|
|responsible_for_future_candidates|Boolean|no|false|The user is responsible for future candidates|
|responsible_for_inactive_candidates|Boolean|no|false|The user is responsible for inactive candidates|

Example:
```
resource "greenhouse_job" "mynewjob" {
  number_openings = 1
  template_job_id = greenhouse_job.mytemplate.id
  anywhere = false
  confidential = false
  custom_fields = {
    field1 = "value1",
    field2 = "value2",
  }
  department_id = greenhouse_department.mydepartment.id
  hiring_team {
    name = "recruiters"
    members {
      user_id = greenhouse_user.recruiter1.id
      responsible_for_active_candidates = true
      responsible_for_future_candidates = true
      responsible_for_inactive_candidates = false
    }
    members {
      user_id = greenhouse_user.recruiter2.id
      responsible_for_inactive_candidates = true
    }
  }
  hiring_team {
    name = "hiring_managers"
    members {
      user_id = greenhouse_user.bigbossperson.id
      responsible_for_active_candidates = true
      responsible_for_future_candidates = true
      responsible_for_inactive_candidates = true
    }
  }
  how_to_sell_this_job = "Say cool things.\nSay nice things.\n"
  job_name = "My cool job"
  job_post_name = "My cool job"
  notes = "Notes go here."
  office_ids =
    greenhouse_office.office1.id,
    greenhouse_office.office2.id,
  opening_ids = [
    greenhouse_opening.myjobopening.id,
  ]
  requisition_id = "abc123"
  team_and_responsibilities = "You work with a team of cool people doing cool things.\nYou are expected to create 500 widgets each day."
}
```

