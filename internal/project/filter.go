package project

type Filter struct {
	Names   []string
	Exclude []ComponentType
}

func ApplyFilter(p *Project, f Filter) *Project {
	if p == nil {
		return nil
	}

	excluded := make(map[ComponentType]struct{}, len(f.Exclude))

	for _, t := range f.Exclude {
		excluded[t] = struct{}{}
	}

	names := make(map[string]struct{}, len(f.Names))

	for _, n := range f.Names {
		names[n] = struct{}{}
	}

	result := &Project{
		Metadata:   p.Metadata,
		Components: make([]Component, 0, len(p.Components)),
	}

	for _, c := range p.Components {

		if _, ok := excluded[c.Type]; ok {
			continue
		}

		if len(names) > 0 {
			if _, ok := names[c.Name]; !ok {
				continue
			}
		}

		result.Components = append(result.Components, c)
	}

	return result
}
