resource "jira_issue" "example" {
  issue_type  = "Task"
  project_key = "PROJ"
  summary     = "Created using Terraform"

  // parent issue is optional
  parent = "PROJ-1"

  // description is optional  
  description = "This is a test issue"

  // (optional) Instead of deleting the issue, perform this transition 
  delete_transition = 21

  // (optional) Make sure, the issue is in the desired state
  // using state_transition
  state            = 10000
  state_transition = 31
}

data "jira_field" "timespent" {
  name = "timespent"
}

resource "jira_issue" "custom_fields_example" {
  issue_type = "Task"
  summary    = "Also Created using Terraform"
  fields = {
    (jira_field.timespent.id) = "30m"
  }
  project_key = "PROJ"
}

