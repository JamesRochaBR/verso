package render

import "testing"

func TestMarkdownRendererRegistered(t *testing.T) {

	r, ok := Get("markdown")

	if !ok {
		t.Fatal("renderer not registered")
	}

	if r.Name() != "markdown" {
		t.Fatal("unexpected renderer")
	}
}
