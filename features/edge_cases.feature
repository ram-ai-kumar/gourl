Feature: Edge Cases
  As a developer
  I want to verify that the application handles edge cases properly
  So that I can ensure reliability in unusual situations

  Background:
    Given I have a clean project directory

  Scenario: Maximum length environment name
    When I run "gourl set env1234567890123456789012345678901234567890123456789012345678901234567890 https://example.com"
    Then the command should succeed
    And the URL should be saved for the long environment name

  Scenario: Single character environment name
    When I run "gourl set a https://example.com"
    Then the command should succeed
    And the URL should be saved for "a"

  Scenario: Environment name with only numbers
    When I run "gourl set 123 https://example.com"
    Then the command should succeed
    And the URL should be saved for "123"

  Scenario: Environment name with mixed case
    When I run "gourl set PrOdUcTiOn https://example.com"
    Then the command should succeed
    And the URL should be saved for "PrOdUcTiOn"

  Scenario: URL with query parameters
    When I run "gourl set production https://example.com?param1=value1&param2=value2"
    Then the command should succeed
    And the URL should be saved with query parameters

  Scenario: URL with fragment identifier
    When I run "gourl set production https://example.com#section1"
    Then the command should succeed
    And the URL should be saved with fragment

  Scenario: URL with authentication
    When I run "gourl set production https://user:pass@example.com"
    Then the command should succeed
    And the URL should be saved with credentials

  Scenario: URL with port number
    When I run "gourl set production https://example.com:8443"
    Then the command should succeed
    And the URL should be saved with port

  Scenario: IPv6 URL
    When I run "gourl set production https://[2001:db8::1]:8443"
    Then the command should succeed
    And the URL should be saved for "production"

  Scenario: Localhost with different protocols
    When I run "gourl set local ftp://localhost:21"
    Then the command should succeed
    And the URL should be saved for "local"

  Scenario: URL with international domain
    When I run "gourl set production https://例子.测试"
    Then the command should succeed
    And the URL should be saved for "production"

  Scenario: URL with trailing slash
    When I run "gourl set production https://example.com/"
    Then the command should succeed
    And the URL should be saved with trailing slash

  Scenario: URL without protocol
    When I run "gourl set production example.com"
    Then the command should succeed
    And the URL should be saved as "example.com"

  Scenario: Multiple consecutive slashes in URL
    When I run "gourl set production https:///example.com"
    Then the command should succeed
    And the URL should be saved as is

  Scenario: URL with encoded characters
    When I run "gourl set production https://example.com/path%20with%20spaces"
    Then the command should succeed
    And the URL should be saved with encoded characters

  Scenario: Empty URL after trimming
    When I run "gourl set production   "
    Then the command should succeed
    And the URL should be saved for "production"

  Scenario: Environment name with leading/trailing spaces
    When I run "gourl set  production  https://example.com  "
    Then the command should succeed
    And the URL should be saved for "production"

  Scenario: URL with consecutive dots
    When I run "gourl set production https://example...com"
    Then the command should succeed
    And the URL should be saved as is

  Scenario: Very long URL
    When I run "gourl set production https://example.com/$(python -c 'print("a" * 1000)')"
    Then the command should succeed
    And the URL should be saved for "production"

  Scenario: Configuration file with BOM
    Given I create a config file with UTF-8 BOM
    And I add valid JSON content
    When I run "gourl list"
    Then the command should succeed
    And I should see the configured URLs

  Scenario: Configuration file with wrong encoding
    Given I create a config file with UTF-16 encoding
    And I add valid JSON content
    When I run "gourl list"
    Then the command should handle the error gracefully

  Scenario: Rapid successive operations
    When I run "gourl set test1 https://example1.com"
    And I run "gourl set test2 https://example2.com"
    And I run "gourl set test3 https://example3.com"
    And I run "gourl list"
    Then the command should succeed
    And I should see all three environments

  Scenario: Concurrent access simulation
    Given I have saved multiple URLs
    When I simulate concurrent file access
    And I run "gourl list"
    Then the command should succeed
    And I should see consistent results

  Scenario: Configuration file during write interruption
    Given I start writing to config file
    When I simulate interruption during write
    And I run "gourl list"
    Then the command should handle the partial file gracefully

  Scenario: Environment name with dots
    When I run "gourl set env.name https://example.com"
    Then the command should succeed
    And the URL should be saved for "env.name"

  Scenario: URL with at symbol in path
    When I run "gourl set production https://example.com/path@file"
    Then the command should succeed
    And the URL should be saved for "production"

  Scenario: Mixed protocol case
    When I run "gourl set production HTTP://example.com"
    Then the command should succeed
    And the URL should be saved as "HTTP://example.com"

  Scenario: URL with bracket notation
    When I run "gourl set production https://example.com/path[1].txt"
    Then the command should succeed
    And the URL should be saved for "production"
