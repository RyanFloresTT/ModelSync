package models

// Config structure
type Config struct {
	WatchFile string   `json:"watchFile"`
	Targets   []Target `json:"targets"`
}

type Target struct {
	Language     string `json:"language"`
	Output       string `json:"output"`
	TemplatePath string `json:"templatePath"`
}
