Feature: Negative Tests
  As a developer
  I want to verify that the application handles invalid inputs gracefully
  So that I can ensure robustness and proper error handling

  Background:
    Given I have a clean project directory

  Scenario: Invalid command names
    When I run "gourl invalidcommand"
    Then the command should fail
    And I should see an error about missing URL for "invalidcommand"

  Scenario: Empty environment name
    When I run "gourl set  https://example.com"
    Then the command should fail
    And I should see usage information for set command

  Scenario: Empty URL
    When I run "gourl set production"
    Then the command should fail
    And I should see usage information for set command

  Scenario: Invalid URL format
    When I run "gourl set production not-a-url"
    Then the command should succeed
    And the URL should be saved for "production"
    When I run "gourl production"
    Then the command should attempt to open "not-a-url"

  Scenario: Extremely long environment name
    When I run "gourl set verylongenvironmentnamethatmightcauseissues https://example.com"
    Then the command should succeed
    And the URL should be saved for "verylongenvironmentnamethatmightcauseissues"

  Scenario: Special characters in environment name
    When I run "gourl set env@#$% https://example.com"
    Then the command should succeed
    And the URL should be saved for "env@#$%"

  Scenario: Unicode characters in environment name
    When I run "gourl set 生产环境 https://example.com"
    Then the command should succeed
    And the URL should be saved for "生产环境"

  Scenario: Null byte in environment name
    When I run "gourl set test\x00env https://example.com"
    Then the command should succeed
    And the URL should be saved for "test"

  Scenario: Newline characters in input
    When I run "gourl set production\nhttps://example.com"
    Then the command should fail
    And I should see usage information for set command

  Scenario: Tab characters in input
    When I run "gourl set\tproduction https://example.com"
    Then the command should fail
    And I should see usage information for set command

  Scenario: URL with JavaScript protocol
    When I run "gourl set production javascript:alert('xss')"
    Then the command should succeed
    And the URL should be saved for "production"

  Scenario: URL with data protocol
    When I run "gourl set production data:text/plain;base64,SGVsbG8gV29ybGQ="
    Then the command should succeed
    And the URL should be saved for "production"

  Scenario: URL with file protocol
    When I run "gourl set production file:///etc/passwd"
    Then the command should succeed
    And the URL should be saved for "production"

  Scenario: Non-existent environment access
    When I run "gourl nonexistentenv"
    Then the command should fail
    And I should see an error about missing URL

  Scenario: Access after deleting config file
    Given I have saved "production" as "https://example.com"
    And .cache/gourls.json exists
    When I delete the config file
    And I run "gourl production"
    Then the command should fail
    And I should see an error about missing URL

  Scenario: Multiple arguments for set command
    When I run "gourl set production https://example.com extra-arg"
    Then the command should succeed
    And the URL should be saved for "production"

  Scenario: Environment name with spaces
    When I run "gourl set production env https://example.com"
    Then the command should succeed
    And the URL should be saved for "production"

  Scenario: Empty configuration directory
    Given .cache directory exists but is empty
    When I run "gourl list"
    Then the command should succeed
    And I should see "No URLs configured"

  Scenario: Permission denied on config directory
    Given I create .cache directory with no permissions
    When I run "gourl set production https://example.com"
    Then the command should fail

  Scenario: Disk full scenario
    Given I simulate disk full condition
    When I run "gourl set production https://example.com"
    Then the command should fail
