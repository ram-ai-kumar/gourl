Feature: Smoke Tests
  As a developer
  I want to verify that the basic functionality works out of the box
  So that I can ensure the application is properly installed and functional

  Background:
    Given I have a clean project directory

  Scenario: Basic installation verification
    When I run "gourl help"
    Then the command should succeed
    And I should see usage information
    And I should see available commands
    And I should see environment aliases

  Scenario: Version verification
    When I run "gourl version"
    Then the command should succeed
    And I should see version information

  Scenario: Basic workflow end-to-end
    When I run "gourl set production https://prod.example.com"
    Then the command should succeed
    And the URL should be saved for "production"
    When I run "gourl list"
    Then the command should succeed
    And I should see "production" in the output
    And I should see "https://prod.example.com" in the output

  Scenario: Default behavior with no configuration
    When I run "gourl"
    Then the command should succeed
    And I should see help message
    And I should not see an error

  Scenario: Configuration file creation
    When I run "gourl set dev http://localhost:3000"
    Then the command should succeed
    And .cache/gourls.json should exist
    And the configuration file should be valid JSON

  Scenario: Multiple environments setup
    When I run "gourl set production https://prod.example.com"
    And I run "gourl set staging https://staging.example.com"
    And I run "gourl set dev http://localhost:3000"
    Then the command should succeed
    When I run "gourl list"
    Then I should see "production" in the output
    And I should see "staging" in the output
    And I should see "dev" in the output
