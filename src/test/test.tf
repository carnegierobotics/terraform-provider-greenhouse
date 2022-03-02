resource "department" "engineering" {
  name = "engineering"
}
resource "department" "software_engineering" {
  name      = "software_engineering"
  # When the department.engineering resource gets created, it will be associated with an ID (and other data). We can reference it here.
  parent_id = department.engineering.id
}
resource "user" "arubin" {
  first_name        = "Adam"
  last_name         = "Rubin"
  email             = "arubin@carnegierobotics.com"
  # Default for send_email_invite is false anyway, but for clarity I included it here.
  send_email_invite = false
}
resource "user" "dfoell" {
  first_name        = "Deborah"
  last_name         = "Foell"
  email             = "dfoell@carnegierobotics.com"
  send_email_invite = false
}
resource "hiring_team" "software" { 
  # When the job.software_engineer resource gets created, it will be associated with an ID (and other data). We can reference it here.
  job_id = job.software_engineer.id
  group {
    name = "people_who_dont_recruit"
    user {
      # When the user.arubin resource gets created, it will be associated with an ID (and other data). We can reference it here.
      id                                  = user.arubin.id
      responsible_for_future_candidates   = false
      responsible_for_active_candidates   = false
      responsible_for_inactive_candidates = false
    }
  }
  group {
    name = "people_who_recruit"
    user {
      id                                  = user.dfoell.id
      responsible_for_future_candidates   = true
      responsible_for_active_candidates   = true
      responsible_for_inactive_candidates = true
    }
  }
}
# You would need to create an "engineer" job template (I don't think it can be done via API); let's pretend it has ID 12345.
resource "job" "software_engineer" { 
  template_job_id    = 12345
  number_of_openings = 1
  job_post_name      = "software_engineer"
  job_name           = "minibot_software_engineer"
  department_id      = department.software_engineering.id
  hiring_team        = hiring_team.software
}
