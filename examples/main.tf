resource "jira_issue" "example_epic" {
  assignee = "anubhavmishra"
  reporter = "anubhavmishra"

  issue_type = "Epic"

  // description is optional
  description = "This is an epic description"
  summary     = "This is an epic summary"

  labels = ["one", "two", "buckle-my-shoe"]

  project_key = "PROJ"
}

resource "jira_issue" "example" {
  assignee = "anubhavmishra"
  reporter = "anubhavmishra"

  issue_type = "Task"

  // description is optional
  description = "This is a test issue that's part of an epic"
  summary     = "Created using Terraform"
  labels      = ["label1", "label2"]
  parent      = resource.jira_issue.example_epic.issue_key
  project_key = "PROJ"
}

