package project

import "testing"

func sampleProject() *Project {
	return &Project{
		Metadata: Metadata{
			Name: "demo",
		},
		Components: []Component{
			{
				Name: "reviewer",
				Type: ComponentSkill,
			},
			{
				Name: "architect",
				Type: ComponentSkill,
			},
			{
				Name: "project",
				Type: ComponentMemory,
			},
		},
	}
}

func TestApplyFilterWithoutFilters(t *testing.T) {
	p := sampleProject()

	got := ApplyFilter(p, Filter{})

	if len(got.Components) != 3 {
		t.Fatalf("expected 3 components, got %d", len(got.Components))
	}
}

func TestApplyFilterByName(t *testing.T) {
	p := sampleProject()

	got := ApplyFilter(p, Filter{
		Names: []string{"reviewer"},
	})

	if len(got.Components) != 1 {
		t.Fatalf("expected 1 component")
	}

	if got.Components[0].Name != "reviewer" {
		t.Fatalf("unexpected component")
	}
}

func TestApplyFilterExcludeType(t *testing.T) {
	p := sampleProject()

	got := ApplyFilter(p, Filter{
		Exclude: []ComponentType{
			ComponentMemory,
		},
	})

	if len(got.Components) != 2 {
		t.Fatalf("expected 2 components")
	}

	for _, c := range got.Components {
		if c.Type == ComponentMemory {
			t.Fatalf("memory should have been excluded")
		}
	}
}

func TestApplyFilterDoesNotMutateOriginal(t *testing.T) {
	p := sampleProject()

	_ = ApplyFilter(p, Filter{
		Names: []string{"reviewer"},
	})

	if len(p.Components) != 3 {
		t.Fatalf("original project was modified")
	}
}
