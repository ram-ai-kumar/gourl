Feature: Environment Aliases
  As a developer
  I want to use shorthand aliases for common environments
  So that I can type less and work faster

  Background:
    Given I have a clean project directory

  Scenario: Production aliases
    When I save URL using "prod" alias
    Then it should be saved as "production"
    When I run "gourl prod"
    Then it should open the production URL
    When I run "gourl p"
    Then it should open the production URL
    When I run "gourl live"
    Then it should open the production URL

  Scenario: Staging aliases
    When I save URL using "stg" alias
    Then it should be saved as "staging"
    When I run "gourl stg"
    Then it should open the staging URL
    When I run "gourl stage"
    Then it should open the staging URL

  Scenario: Development aliases
    When I save URL using "d" alias
    Then it should be saved as "dev"
    When I run "gourl d"
    Then it should open the dev URL
    When I run "gourl local"
    Then it should open the dev URL

  Scenario: Custom environment without alias
    When I save URL for "custom" environment
    Then it should be saved as "custom"
    When I run "gourl custom"
    Then it should open the custom URL
