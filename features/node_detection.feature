Feature: Node.js Project Detection
  In order to configure Node.js projects easily
  As a developer
  I want the interactive setup to recognize Node.js markers

  Scenario: Detect Node.js project via package.json
    Given a project with a "package.json" file
    When I run "gourl -i" and provide inputs:
      | http://localhost:3000 |
      | https://stg.example.com |
      | https://prod.example.com |
    Then output should contain "Detected Node.js project"
    And output should contain "Suggested dev port: 3000"

  Scenario: Detect Node.js project via yarn.lock
    Given a project with a "yarn.lock" file
    When I run "gourl -i" and provide inputs:
      | http://localhost:3000 |
      | https://stg.example.com |
      | https://prod.example.com |
    Then output should contain "Detected Node.js (Yarn) project"

  Scenario: Detect Node.js project via pnpm-lock.yaml
    Given a project with a "pnpm-lock.yaml" file
    When I run "gourl -i" and provide inputs:
      | http://localhost:3000 |
      | https://stg.example.com |
      | https://prod.example.com |
    Then output should contain "Detected Node.js (pnpm) project"
