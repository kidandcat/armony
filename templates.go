package armony

import (
	"html/template"
)

// Templates : *template.Template
var Templates *template.Template

// ParseTemplates : Parse templates from configured folders
func ParseTemplates(templateFolders []string) {
	Templates = &template.Template{}
	empty := true
	for _, v := range templateFolders {
		if empty {
			Templates = template.Must(template.ParseGlob(v + "/*.html"))
			empty = false
		} else {
			Templates = template.Must(Templates.ParseGlob(v + "/*.html"))
		}
	}
}
