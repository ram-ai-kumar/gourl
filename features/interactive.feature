Feature: Interactive Setup
  In order to configure URLs easily
  As a developer
  I want a guided interactive setup mode

  Scenario: Interactive setup in a Go project suggests default port
    Given a project with a "go.mod" file
    When I run "gourl -i" and provide inputs:
      | http://localhost:8080 |
      | https://stg.example.com |
      | https://prod.example.com |
    Then the ".cache/gourls.json" file should exist
    And the URL should be saved for "dev"
    And output should contain "Detected Go project"
    And output should contain "Suggested dev port: 8080"
