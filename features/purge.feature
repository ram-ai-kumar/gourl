Feature: Purge Functionality
  In order to uninstall gourl easily
  As a developer
  I want to be able to remove the binary using the --purge command

  Scenario: Purge the binary with force
    Given a project with a compiled gourl binary
    When I run "gourl --purge --force"
    Then the binary should be removed from the system
    And the ".cache/gourls.json" file should still exist
