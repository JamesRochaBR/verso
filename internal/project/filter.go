package project

type Filter struct {
	Skills  []string
	Memory  []string
	Exclude []ComponentType
}

func ApplyFilter(p *Project, f Filter) *Project {
	return p
}
