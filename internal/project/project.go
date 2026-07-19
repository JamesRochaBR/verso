package project

type Project struct {
	Metadata   Metadata    `json:"metadata"`
	Components []Component `json:"components"`
}
