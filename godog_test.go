package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/cucumber/godog"
)

const (
	testDir    = "/tmp/gourl-test"
	testConfigPath = ".cache/gourls.json"
)

var (
	lastOutput     string
	lastExitCode   int
	lastError      error
)

type testSuite struct{}

func (t *testSuite) InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
		// Set test mode to stub URL opening
		os.Setenv("GOURL_TEST_MODE", "1")

		// Create test directory
		if err := os.RemoveAll(testDir); err != nil {
			fmt.Printf("Failed to clean test directory: %v\n", err)
			os.Exit(1)
		}

		if err := os.MkdirAll(testDir, 0755); err != nil {
			fmt.Printf("Failed to create test directory: %v\n", err)
			os.Exit(1)
		}

		// Change to test directory
		if err := os.Chdir(testDir); err != nil {
			fmt.Printf("Failed to change to test directory: %v\n", err)
			os.Exit(1)
		}
	})
}

func (t *testSuite) InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Given(`^I have a clean project directory$`, func() error {
		if _, err := os.Stat(testConfigPath); err == nil {
			return fmt.Errorf("config file should not exist")
		}
		return nil
	})

	ctx.Given(`^no \.cache/gourls\.json file exists$`, func() error {
		if _, err := os.Stat(testConfigPath); err == nil {
			return fmt.Errorf("config file should not exist")
		}
		return nil
	})

	ctx.Given(`^I have saved URLs:$`, func(table *godog.Table) error {
		for _, row := range table.Rows[1:] { // Skip header
			env := row.Cells[0].Value
			url := row.Cells[1].Value
			t.runGourlCommand("set", env, url)
		}
		return nil
	})

	ctx.When(`^I run "gourl list"$`, func(args string) error {
		argParts := strings.Fields(args)
		return t.runGourlCommand(argParts...)
	})

	ctx.Then(`^I should see configured URLs$`, func() error {
		if !strings.Contains(lastOutput, "Configured URLs:") {
			return fmt.Errorf("output should contain configured URLs")
		}
		return nil
	})

	ctx.Then(`^output should contain "([^"]*)"$`, func(text string) error {
		if !strings.Contains(lastOutput, text) {
			return fmt.Errorf("output should contain %s", text)
		}
		return nil
	})

	ctx.Given(`^I have saved "([^"]*)" as "([^"]*)"$`, func(env, url string) error {
		return t.runGourlCommand("set", env, url)
	})

	ctx.When(`^I run "gourl ([^"]*)"$`, func(args string) error {
		argParts := strings.Fields(args)
		return t.runGourlCommand(argParts...)
	})

	ctx.Then(`^it should open ([^"]*) URL$`, func(envType string) error {
		if lastExitCode != 0 {
			return fmt.Errorf("command should succeed when opening %s URL", envType)
		}
		return nil
	})

	ctx.Then(`^the command should succeed$`, func() error {
		if lastExitCode != 0 {
			return fmt.Errorf("command should succeed")
		}
		return nil
	})

	ctx.Then(`^the command should fail$`, func() error {
		if lastExitCode == 0 {
			return fmt.Errorf("command should have failed")
		}
		return nil
	})

	ctx.Then(`^the URL should be saved for "([^"]*)"$`, func(env string) error {
		data, err := os.ReadFile(testConfigPath)
		if err != nil {
			return fmt.Errorf("config file should exist")
		}

		var config map[string]string
		err = json.Unmarshal(data, &config)
		if err != nil {
			return fmt.Errorf("config should be valid JSON")
		}

		if _, ok := config[env]; !ok {
			return fmt.Errorf("URL not found for environment %s", env)
		}

		return nil
	})

	ctx.Then(`^\.cache/gourls\.json should exist$`, func() error {
		if _, err := os.Stat(testConfigPath); err != nil {
			return fmt.Errorf("config file should exist")
		}
		return nil
	})

	// Environment aliases
	ctx.Given(`^I save URL using "([^"]*)" alias$`, func(alias string) error {
		url := fmt.Sprintf("https://%s.example.com", alias)
		return t.runGourlCommand("set", alias, url)
	})

	ctx.Then(`^it should be saved as "([^"]*)"$`, func(env string) error {
		return t.theURLShouldBeSavedFor(env)
	})

	ctx.When(`^I run "gourl ([^"]*)"$`, func(args string) error {
		argParts := strings.Fields(args)
		return t.runGourlCommand(argParts...)
	})

	ctx.Then(`^it should open the ([^"]*) URL$`, func(env string) error {
		if lastExitCode != 0 {
			return fmt.Errorf("command should succeed when opening %s URL", env)
		}
		return nil
	})

	// Edge cases and error handling
	ctx.Given(`^\.cache/gourls\.json exists but is empty$`, func() error {
		err := os.MkdirAll(filepath.Dir(testConfigPath), 0755)
		if err != nil {
			return err
		}
		return os.WriteFile(testConfigPath, []byte{}, 0644)
	})

	ctx.Given(`^\.cache/gourls\.json contains invalid JSON$`, func() error {
		err := os.MkdirAll(filepath.Dir(testConfigPath), 0755)
		if err != nil {
			return err
		}
		return os.WriteFile(testConfigPath, []byte("{ invalid json"), 0644)
	})

	ctx.Then(`^the command should handle the error gracefully$`, func() error {
		if !strings.Contains(lastOutput, "No URLs configured") {
			return fmt.Errorf("output should handle error gracefully")
		}
		return nil
	})

	ctx.Given(`^\.gitignore does not contain "\.cache/"$`, func() error {
		gitignorePath := filepath.Join(testDir, ".gitignore")
		return os.WriteFile(gitignorePath, []byte("node_modules/\n"), 0644)
	})

	ctx.Given(`^\.gitignore contains "\.cache/"$`, func() error {
		gitignorePath := filepath.Join(testDir, ".gitignore")
		return os.WriteFile(gitignorePath, []byte(".cache/\nnode_modules/\n"), 0644)
	})

	// Smoke tests
	ctx.Step(`^I should see usage information for set command$`, func() error {
		if !strings.Contains(lastOutput, "Usage: gourl set <env> <url>") {
			return fmt.Errorf("output should contain usage information")
		}
		return nil
	})

	ctx.Step(`^I should see available commands$`, func() error {
		if !strings.Contains(lastOutput, "Available commands:") {
			return fmt.Errorf("output should contain available commands")
		}
		return nil
	})

	ctx.Step(`^I should see environment aliases$`, func() error {
		if !strings.Contains(lastOutput, "Environment aliases:") {
			return fmt.Errorf("output should contain environment aliases")
		}
		return nil
	})

	ctx.Step(`^I should see version information$`, func() error {
		if !strings.Contains(lastOutput, "gourl version") {
			return fmt.Errorf("output should contain version information")
		}
		return nil
	})

	ctx.Step(`^I should see help message$`, func() error {
		if !strings.Contains(lastOutput, "Usage:") {
			return fmt.Errorf("output should contain help message")
		}
		return nil
	})

	ctx.Step(`^I should not see an error$`, func() error {
		if strings.Contains(strings.ToLower(lastOutput), "error") {
			return fmt.Errorf("output should not contain error")
		}
		return nil
	})

	ctx.Step(`^I should see an error about missing URL$`, func() error {
		if !strings.Contains(lastOutput, "No URL found") {
			return fmt.Errorf("output should contain missing URL error")
		}
		return nil
	})

	ctx.Step(`^I should see an error about missing URL for "([^"]*)"$`, func(env string) error {
		expected := fmt.Sprintf("No URL found for '%s'", env)
		if !strings.Contains(lastOutput, expected) {
			return fmt.Errorf("output should contain error about missing URL for %s", env)
		}
		return nil
	})

	// Security tests
	ctx.Step(`^I should see a security warning about untrusted URL$`, func() error {
		if !strings.Contains(lastOutput, "untrusted") {
			return fmt.Errorf("output should contain security warning about untrusted URL")
		}
		return nil
	})

	ctx.Step(`^I should see a warning about privileged access$`, func() error {
		if !strings.Contains(lastOutput, "privileged") {
			return fmt.Errorf("output should contain warning about privileged access")
		}
		return nil
	})

	ctx.Step(`^I should see verification message$`, func() error {
		if !strings.Contains(lastOutput, "verification") {
			return fmt.Errorf("output should contain verification message")
		}
		return nil
	})

	ctx.Step(`^I should see a warning about internal network access$`, func() error {
		if !strings.Contains(lastOutput, "internal") {
			return fmt.Errorf("output should contain warning about internal network access")
		}
		return nil
	})

	ctx.Step(`^I should see a warning about external network access$`, func() error {
		if !strings.Contains(lastOutput, "external") {
			return fmt.Errorf("output should contain warning about external network access")
		}
		return nil
	})

	ctx.Step(`^I should see a security warning about mobile device$`, func() error {
		if !strings.Contains(lastOutput, "mobile") {
			return fmt.Errorf("output should contain security warning about mobile device")
		}
		return nil
	})

	// ZTA tests
	ctx.Step(`^I should see session information$`, func() error {
		if !strings.Contains(lastOutput, "session") {
			return fmt.Errorf("output should contain session information")
		}
		return nil
	})

	ctx.Step(`^I should see context validation message$`, func() error {
		if !strings.Contains(lastOutput, "context") {
			return fmt.Errorf("output should contain context validation message")
		}
		return nil
	})

	ctx.Step(`^I should see real-time verification status$`, func() error {
		if !strings.Contains(lastOutput, "real-time") {
			return fmt.Errorf("output should contain real-time verification status")
		}
		return nil
	})

	ctx.Step(`^I should see authentication requirement$`, func() error {
		if !strings.Contains(lastOutput, "authentication") {
			return fmt.Errorf("output should contain authentication requirement")
		}
		return nil
	})

	ctx.Step(`^I should see resource isolation message$`, func() error {
		if !strings.Contains(lastOutput, "isolation") {
			return fmt.Errorf("output should contain resource isolation message")
		}
		return nil
	})

	// Compliance tests
	ctx.Step(`^only necessary data should be stored$`, func() error {
		data, err := os.ReadFile(testConfigPath)
		if err != nil {
			return fmt.Errorf("config file should exist")
		}
		var config map[string]string
		err = json.Unmarshal(data, &config)
		if len(config) > 1 { // Only test URL should be stored
			return fmt.Errorf("only necessary data should be stored")
		}
		return nil
	})

	ctx.Step(`^no personal data should be included without consent$`, func() error {
		data, err := os.ReadFile(testConfigPath)
		if err != nil {
			return fmt.Errorf("config file should exist")
		}
		var config map[string]string
		err = json.Unmarshal(data, &config)
		// Check for personal data patterns
		personalPatterns := []string{"email", "name", "phone", "address"}
		for _, pattern := range personalPatterns {
			for _, value := range config {
				if strings.Contains(strings.ToLower(value), pattern) {
					return fmt.Errorf("no personal data should be included without consent")
				}
			}
		}
		return nil
	})

	ctx.Step(`^access should be logged with timestamp$`, func() error {
		if !strings.Contains(lastOutput, "timestamp") {
			return fmt.Errorf("access should be logged with timestamp")
		}
		return nil
	})

	ctx.Step(`^log should contain user identifier$`, func() error {
		if !strings.Contains(lastOutput, "user") {
			return fmt.Errorf("log should contain user identifier")
		}
		return nil
	})

	ctx.Step(`^log should not contain sensitive data$`, func() error {
		data, err := os.ReadFile(testConfigPath)
		if err != nil {
			return fmt.Errorf("config file should exist")
		}
		var config map[string]string
		err = json.Unmarshal(data, &config)
		// Check for sensitive data patterns
		sensitivePatterns := []string{"password", "token", "secret", "key"}
		for _, pattern := range sensitivePatterns {
			for _, value := range config {
				if strings.Contains(strings.ToLower(value), pattern) {
					return fmt.Errorf("log should not contain sensitive data")
				}
			}
		}
		return nil
	})

	ctx.Step(`^file should have appropriate permissions$`, func() error {
		// Check file permissions
		fileInfo, err := os.Stat(testConfigPath)
		if err != nil {
			return fmt.Errorf("config file should exist")
		}
		// File should not be world-readable
		if fileInfo.Mode().Perm()&0007 != 0600 { // rw-------
			return fmt.Errorf("file should have appropriate permissions")
		}
		return nil
	})

	ctx.Step(`^credentials should be stored securely$`, func() error {
		data, err := os.ReadFile(testConfigPath)
		if err != nil {
			return fmt.Errorf("config file should exist")
		}
		var config map[string]string
		err = json.Unmarshal(data, &config)
		// Check if credentials are stored (this is a simplified check)
		for _, value := range config {
			if strings.Contains(value, "user:pass") {
				return fmt.Errorf("credentials should be stored securely")
			}
		}
		return nil
	})

	// Additional step definitions for comprehensive test coverage
	ctx.Step(`^I should see international character support$`, func() error {
		if !strings.Contains(lastOutput, "生产环境") {
			return fmt.Errorf("output should contain international character support")
		}
		return nil
	})

	ctx.Step(`^I should see financial compliance warning$`, func() error {
		if !strings.Contains(lastOutput, "financial") {
			return fmt.Errorf("output should contain financial compliance warning")
		}
		return nil
	})

	ctx.Step(`^I should see healthcare compliance warning$`, func() error {
		if !strings.Contains(lastOutput, "healthcare") {
			return fmt.Errorf("output should contain healthcare compliance warning")
		}
		return nil
	})

	ctx.Step(`^I should see export control notice$`, func() error {
		if !strings.Contains(lastOutput, "export") {
			return fmt.Errorf("output should contain export control notice")
		}
		return nil
	})

	ctx.Step(`^I should see data sovereignty notice$`, func() error {
		if !strings.Contains(lastOutput, "sovereignty") {
			return fmt.Errorf("output should contain data sovereignty notice")
		}
		return nil
	})

	ctx.Step(`^I should see incident response protocol$`, func() error {
		if !strings.Contains(lastOutput, "incident") {
			return fmt.Errorf("output should contain incident response protocol")
		}
		return nil
	})

	ctx.Step(`^I should see vulnerability disclosure notice$`, func() error {
		if !strings.Contains(lastOutput, "vulnerability") {
			return fmt.Errorf("output should contain vulnerability disclosure notice")
		}
		return nil
	})

	ctx.Step(`^I should see SLA compliance notice$`, func() error {
		if !strings.Contains(lastOutput, "SLA") {
			return fmt.Errorf("output should contain SLA compliance notice")
		}
		return nil
	})

	ctx.Step(`^I should see service level requirements$`, func() error {
		if !strings.Contains(lastOutput, "service level") {
			return fmt.Errorf("output should contain service level requirements")
		}
		return nil
	})

	// Then steps
	ctx.Then(`^the configuration file should be valid JSON$`, func() error {
		data, err := os.ReadFile(testConfigPath)
		if err != nil {
			return fmt.Errorf("config file should exist")
		}
		var config map[string]string
		err = json.Unmarshal(data, &config)
		if err != nil {
			return fmt.Errorf("config should be valid JSON")
		}
		return nil
	})
}

func (t *testSuite) runGourlCommand(args ...string) error {
	binaryPath := filepath.Join(testDir, "gourl")
	if runtime.GOOS == "windows" {
		binaryPath += ".exe"
	}

	// Build binary
	buildCmd := exec.Command("go", "build", "-o", binaryPath, ".")
	buildCmd.Dir = "/Users/ram/Work/code/dev-stack/gourl"
	output, err := buildCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to build gourl: %v\nBuild output: %s", err, string(output))
	}

	// Run command with test mode environment
	cmd := exec.Command(binaryPath, args...)
	cmd.Dir = testDir
	cmd.Env = append(os.Environ(), "GOURL_TEST_MODE=1")
	output, err = cmd.CombinedOutput()
	lastOutput = string(output)
	if err != nil {
		lastExitCode = 1
		lastError = err
	} else {
		lastExitCode = 0
	}

	return nil
}

func (t *testSuite) theURLShouldBeSavedFor(env string) error {
	data, err := os.ReadFile(testConfigPath)
	if err != nil {
		return fmt.Errorf("config file should exist")
	}

	var config map[string]string
	err = json.Unmarshal(data, &config)
	if err != nil {
		return fmt.Errorf("config should be valid JSON")
	}

	if _, ok := config[env]; !ok {
		return fmt.Errorf("URL not found for environment %s", env)
	}

	return nil
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: func(ctx *godog.ScenarioContext) {
			testSuite := &testSuite{}
			testSuite.InitializeScenario(ctx)
		},
		TestSuiteInitializer: func(ctx *godog.TestSuiteContext) {
			testSuite := &testSuite{}
			testSuite.InitializeTestSuite(ctx)
		},
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("failed to run feature tests")
	}
}
