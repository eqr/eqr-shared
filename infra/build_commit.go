package infra

import (
	"fmt"
	"os"
	"strings"
)

type BuildCommit struct {
	Commit   string
	Date     string
	Branch   string
	HasValue bool
}

func (buildCommit *BuildCommit) BuildCommitFromEnv() error {
	if val, err := readFromEnv("COMMIT"); err != nil {
		return fmt.Errorf("cannot read commit from environment: %w", err)
	} else {
		buildCommit.Commit = val
	}

	if val, err := readFromEnv("BRANCH"); err != nil {
		return fmt.Errorf("cannot read branch from environment: %w", err)
	} else {
		buildCommit.Branch = val
	}

	if val, err := readFromEnv("DATE"); err != nil {
		return fmt.Errorf("cannot read date from environment: %w", err)
	} else {
		buildCommit.Date = val
	}

	buildCommit.HasValue = true
	return nil
}

func readFromEnv(field string) (string, error) {
	value := strings.TrimSpace(os.Getenv(field))
	if value == "" {
		return "", fmt.Errorf("%q cannot be read from env", strings.ToLower(field))
	}

	return value, nil
}

func (buildCommit *BuildCommit) String() string {
	if !buildCommit.HasValue {
		return "Build information is unavailable"
	} else {
		return fmt.Sprintf("Built from %s (%s) at %s", buildCommit.Branch, buildCommit.Commit, buildCommit.Date)
	}
}
