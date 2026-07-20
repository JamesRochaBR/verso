package render

import "github.com/james-rocha/verso/internal/project"

type Renderer interface {
	Name() string
	Render(*project.Project) (string, error)
}

var renderers = map[string]Renderer{}

func Register(r Renderer) {
	renderers[r.Name()] = r
}

func Get(name string) (Renderer, bool) {
	r, ok := renderers[name]
	return r, ok
}
