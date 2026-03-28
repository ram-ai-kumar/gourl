Feature: Security Tests
  As a developer
  I want to verify that the application handles security-related inputs safely
  So that I can ensure the application is secure and doesn't expose vulnerabilities

  Background:
    Given I have a clean project directory

  Scenario: XSS attempt in environment name
    When I run "gourl set <script>alert('xss')</script> https://example.com"
    Then the command should succeed
    And the URL should be saved for the environment name

  Scenario: XSS attempt in URL
    When I run "gourl set production https://example.com/<script>alert('xss')</script>"
    Then the command should succeed
    And the URL should be saved for "production"

  Scenario: SQL injection attempt in environment name
    When I run "gourl set prod'; DROP TABLE gourls; -- https://example.com"
    Then the command should succeed
    And the URL should be saved for the environment name

  Scenario: Command injection in URL
    When I run "gourl set production https://example.com; rm -rf /"
    Then the command should succeed
    And the URL should be saved for "production"

  Scenario: Path traversal attempt in URL
    When I run "gourl set production https://example.com/../../../etc/passwd"
    Then the command should succeed
    And the URL should be saved for "production"

  Scenario: Null byte injection attempt
    When I run "gourl set test\x00https://example.com"
    Then the command should succeed
    And the URL should be saved for "test"

  Scenario: Buffer overflow attempt in environment name
    When I run "gourl set $(python -c 'print("A" * 10000)') https://example.com"
    Then the command should succeed
    And the application should not crash

  Scenario: Format string attack in environment name
    When I run "gourl set %n%n%n%n%n%n https://example.com"
    Then the command should succeed
    And the URL should be saved for the environment name

  Scenario: Unicode normalization attack
    When I run "gourl set ﬀ https://example.com"
    Then the command should succeed
    And the URL should be saved for the environment name

  Scenario: HTML entity encoding in URL
    When I run "gourl set production https://example.com/&lt;script&gt;alert('xss')&lt;/script&gt;"
    Then the command should succeed
    And the URL should be saved for "production"

  Scenario: URL with JavaScript pseudo-protocol
    When I run "gourl set production \x00javascript:alert('xss')"
    Then the command should succeed
    And the URL should be saved for "production"

  Scenario: Environment name with control characters
    When I run "gourl set test\x01\x02\x03 https://example.com"
    Then the command should succeed
    And the URL should be saved for the environment name

  Scenario: URL with log4j vulnerability pattern
    When I run "gourl set production https://example.com/${jndi:ldap://example.com/a}"
    Then the command should succeed
    And the URL should be saved for "production"

  Scenario: Shell metacharacters in environment name
    When I run "gourl set `whoami` https://example.com"
    Then the command should succeed
    And the URL should be saved for the environment name

  Scenario: Environment name with pipe character
    When I run "gourl set test|env https://example.com"
    Then the command should succeed
    And the URL should be saved for "test"

  Scenario: URL with local file inclusion attempt
    When I run "gourl set production file:///etc/passwd%00"
    Then the command should succeed
    And the URL should be saved for "production"

  Scenario: URL with SMB protocol attempt
    When I run "gourl set production smb://evil.com/share"
    Then the command should succeed
    And the URL should be saved for "production"

  Scenario: Environment name with newlines
    When I run "gourl set test\nenv https://example.com"
    Then the command should succeed
    And the URL should be saved for "test"

  Scenario: URL with tab characters
    When I run "gourl set production https://example.com\tmalicious"
    Then the command should succeed
    And the URL should be saved for "production"

  Scenario: Environment name with backticks
    When I run "gourl set `rm -rf /` https://example.com"
    Then the command should succeed
    And the URL should be saved for the environment name

  Scenario: URL with environment variable expansion
    When I run "gourl set production https://example.com/$HOME"
    Then the command should succeed
    And the URL should be saved for "production"

  Scenario: URL with command substitution
    When I run "gourl set production https://example.com/$(id)"
    Then the command should succeed
    And the URL should be saved for "production"

  Scenario: URL with LDAP injection pattern
    When I run "gourl set production https://example.com/*%0a"
    Then the command should succeed
    And the URL should be saved for "production"

  Scenario: Environment name with ANSI escape sequences
    When I run "gourl set \x1b[31mtest\x1b[0m https://example.com"
    Then the command should succeed
    And the URL should be saved for the environment name

  Scenario: URL with XML external entity
    When I run "gourl set production https://example.com/&xxe;"
    Then the command should succeed
    And the URL should be saved for "production"
