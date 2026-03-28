Feature: Compliance Tests
  As a developer
  I want to verify that the application meets compliance requirements
  So that I can ensure adherence to standards and regulations

  Background:
    Given I have a clean project directory

  Scenario: GDPR compliance - data minimization
    When I run "gourl set production https://prod.example.com"
    Then the command should succeed
    And the URL should be saved for "production"
    When I inspect the configuration file
    Then only necessary data should be stored
    And no personal data should be included without consent

  Scenario: Data retention policy compliance
    Given I have saved "production" as "https://prod.example.com"
    And I have saved "staging" as "https://staging.example.com"
    When I run "gourl list"
    Then the command should succeed
    And I should see all configured URLs
    When I delete "staging"
    And I run "gourl list"
    Then the command should succeed
    And I should see only "production"
    And the deleted URL should be properly removed

  Scenario: Access logging compliance
    Given I have saved "production" as "https://prod.example.com"
    When I run "gourl production"
    Then it should open the production URL
    And the access should be logged with timestamp
    And the log should contain user identifier
    And the log should not contain sensitive data

  Scenario: Configuration file permissions compliance
    When I run "gourl set production https://prod.example.com"
    Then the command should succeed
    And the configuration file should have appropriate permissions
    And the file should be readable by owner only
    And the file should not be world-readable

  Scenario: URL validation compliance
    When I run "gourl set production https://user:pass@example.com"
    Then the command should succeed
    And the URL should be saved for "production"
    And the credentials should be stored securely
    And I should see a warning about credential storage

  Scenario: Audit trail compliance
    Given I have saved multiple URLs
    When I run "gourl list"
    Then the command should succeed
    And I should see all configured URLs
    When I run "gourl set production https://new-prod.example.com"
    Then the command should succeed
    And the change should be logged
    And I should see audit confirmation

  Scenario: Privacy policy compliance
    Given I have saved "production" as "https://prod.example.com"
    When I run "gourl production"
    Then it should open the production URL
    And no tracking should be performed without consent
    And I should see privacy notice

  Scenario: Security standards compliance
    When I run "gourl set production http://example.com"
    Then the command should succeed
    And the URL should be saved for "production"
    And I should see a warning about insecure protocol
    When I run "gourl set production https://example.com"
    Then the command should succeed
    And I should not see security warnings

  Scenario: Accessibility compliance
    Given I have saved "production" as "https://prod.example.com"
    When I run "gourl production"
    Then it should open the production URL
    And the output should be accessible
    And I should see accessibility information

  Scenario: International compliance
    When I run "gourl set 生产环境 https://china.example.com"
    Then the command should succeed
    And the URL should be saved for "生产环境"
    And I should see international character support

  Scenario: Regulatory compliance - financial
    Given I have saved "payment" as "https://payment.example.com"
    When I run "gourl payment"
    Then it should open the payment URL
    And I should see financial compliance warning
    And the access should require additional verification

  Scenario: Regulatory compliance - healthcare
    Given I have saved "phi" as "https://health.example.com"
    When I run "gourl phi"
    Then it should open the health URL
    And I should see healthcare compliance warning
    And the access should require additional verification

  Scenario: Export control compliance
    Given I have saved "production" as "https://prod.example.com"
    When I run "gourl production"
    Then it should open the production URL
    And I should see export control notice
    And the access should be validated against export regulations

  Scenario: Data sovereignty compliance
    Given I have saved "production" as "https://prod.example.com"
    And current location is in regulated region
    When I run "gourl production"
    Then it should open the production URL
    And I should see data sovereignty notice
    And the data should be stored in compliant region

  Scenario: Incident response compliance
    Given I have saved "production" as "https://prod.example.com"
    When I simulate security incident
    And I run "gourl production"
    Then it should open the production URL
    And I should see incident response protocol
    And the access should be logged for incident tracking

  Scenario: Vulnerability disclosure compliance
    Given I have saved "production" as "https://prod.example.com"
    When I run "gourl production"
    Then it should open the production URL
    And I should see vulnerability disclosure notice
    And the access should be validated for security issues

  Scenario: Service level agreement compliance
    Given I have saved "production" as "https://prod.example.com"
    When I run "gourl production"
    Then it should open the production URL
    And I should see SLA compliance notice
    And the access should meet service level requirements
