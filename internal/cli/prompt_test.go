package cli

import "testing"

func TestPromptCommandRequiresPath(t *testing.T) {

	cmd := PromptCommand{}

	if err := cmd.Run(nil); err == nil {
		t.Fatal("expected error when project path is missing")
	}
}
