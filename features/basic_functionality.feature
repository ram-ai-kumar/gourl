Feature: Basic URL Management Functionality
  As a developer working on a project
  I want to save, list, and open project URLs
  So that I can quickly access different environments

  Background:
    Given I have a clean project directory
    And no .cache/gourls.json file exists

  Scenario: Save a new URL
    When I run "gourl set production https://example.com"
    Then the command should succeed
    And the URL should be saved for "production"
    And .cache/gourls.json should exist

  Scenario: List saved URLs
    Given I have saved URLs:
      | environment | url                    |
      | production  | https://prod.example.com |
      | staging     | https://staging.example.com |
      | dev         | http://localhost:3000    |
    When I run "gourl list"
    Then I should see the configured URLs
    And the output should contain "production"
    And the output should contain "staging"
    And the output should contain "dev"

  Scenario: Open saved URL
    Given I have saved "production" as "https://example.com"
    When I run "gourl production"
    Then the command should attempt to open "https://example.com"

  Scenario: Open non-existent URL
    When I run "gourl nonexistent"
    Then the command should fail
    And I should see an error message about missing URL
    And I should see a suggestion to run "gourl set"

  Scenario: Show help
    When I run "gourl help"
    Then I should see usage information
    And I should see available commands
    And I should see environment aliases

  Scenario: Show version
    When I run "gourl version"
    Then I should see version information
