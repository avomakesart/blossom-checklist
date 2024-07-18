package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func TestCLIExitCodeWhenNotAllItemsChecked(t *testing.T) {
	cmd := exec.Command("./blossom-checklist")
	input := bytes.NewBufferString("n\n")
	cmd.Stdin = input

	output := &bytes.Buffer{}
	cmd.Stdout = output

	err := cmd.Run()
	if err == nil {
		t.Fatalf("expected exit code 1, got %v", err)
	}

	if !strings.Contains(output.String(), "The following items are still unchecked:") {
		t.Fatalf("expected output to contain 'The following items are still unchecked:', got %s", output.String())
	}
}

func TestCLIExitCodeWhenAllItemsChecked(t *testing.T) {
	cmd := exec.Command("./blossom-checklist")
	input := bytes.NewBufferString(strings.Repeat("y\n", 30)) // Adjust the repeat count based on the number of inputs
	cmd.Stdin = input

	output := &bytes.Buffer{}
	cmd.Stdout = output

	err := cmd.Run()
	if err != nil {
		t.Fatalf("expected exit code 0, got %v", err)
	}

	if !strings.Contains(output.String(), "Great! You have completed all checks. You can proceed with your push.") {
		t.Fatalf("expected output to contain 'Great! You have completed all checks. You can proceed with your push.', got %s", output.String())
	}
}

func TestCLIDisplayAllSections(t *testing.T) {
	cmd := exec.Command("./blossom-checklist")
	input := bytes.NewBufferString("n\n")
	cmd.Stdin = input

	output := &bytes.Buffer{}
	cmd.Stdout = output

	err := cmd.Run()
	if err == nil {
		t.Fatalf("expected exit code 1, got %v", err)
	}

	expectedSections := []string{
		"Mobile",
		"Unit test",
		"Cross browser test",
		"Different accounts",
		"Translations",
		"Mock data",
		"Happy flows",
		"After merging",
	}

	for _, section := range expectedSections {
		if !strings.Contains(output.String(), section) {
			t.Fatalf("expected output to contain %s, got %s", section, output.String())
		}
	}
}
