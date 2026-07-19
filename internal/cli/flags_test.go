package cli

import (
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
