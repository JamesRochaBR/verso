package cli

import (
	"github.com/james-rocha/verso/internal/project"
	"reflect"
	"testing"
)

func TestSplitCSV(t *testing.T) {

	got := splitCSV("reviewer,architect,memory")

	want := []string{
		"reviewer",
		"architect",
		"memory",
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatal("unexpected result")
	}
}

func TestSplitCSVTrim(t *testing.T) {

	got := splitCSV(" reviewer , architect ")

	want := []string{
		"reviewer",
		"architect",
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatal("unexpected result")
	}
}

func TestSplitCSVEmpty(t *testing.T) {

	got := splitCSV("")

	if got != nil {
		t.Fatal("expected nil")
	}
}

func TestParsePromptOptions(t *testing.T) {

	opts, err := ParsePromptOptions([]string{
		"--name",
		"reviewer,architect",
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(opts.Filter.Names) != 2 {
		t.Fatal("expected two names")
	}

	if opts.Filter.Names[0] != "reviewer" {
		t.Fatal("unexpected first name")
	}

	if opts.Filter.Names[1] != "architect" {
		t.Fatal("unexpected second name")
	}
}

func TestParsePromptOptionsOutput(t *testing.T) {

	opts, err := ParsePromptOptions([]string{
		"--output",
		"prompt.md",
	})
	if err != nil {
		t.Fatal(err)
	}

	if opts.Output != "prompt.md" {
		t.Fatal("unexpected output")
	}
}

func TestParseComponentTypes(t *testing.T) {

	got := parseComponentTypes("memory,template")

	if len(got) != 2 {
		t.Fatal("expected two component types")
	}

	if got[0] != project.ComponentMemory {
		t.Fatal("unexpected first type")
	}

	if got[1] != project.ComponentTemplate {
		t.Fatal("unexpected second type")
	}
}

func TestParsePromptOptionsExclude(t *testing.T) {

	opts, err := ParsePromptOptions([]string{
		"--exclude",
		"memory",
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(opts.Filter.Exclude) != 1 {
		t.Fatal("expected one excluded type")
	}

	if opts.Filter.Exclude[0] != project.ComponentMemory {
		t.Fatal("unexpected excluded type")
	}
}

func TestParsePromptOptionsUnknownFlag(t *testing.T) {
	_, err := ParsePromptOptions([]string{
		"--invalid",
	})

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestParsePromptOptionsMissingValue(t *testing.T) {
	_, err := ParsePromptOptions([]string{
		"--name",
	})

	if err == nil {
		t.Fatal("expected error")
	}
}
