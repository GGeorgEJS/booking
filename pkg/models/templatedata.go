package models

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[int]int
	FloatMap  map[string]float64
	Data      map[string]interface{}
	CSRToken  string
	Flash     string
	Warning   string
	Error     string
}
