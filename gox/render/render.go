package render

import "html/template"

// The GoX Router will render and serve the full page
// - Value: A struct containing the data you want executed in your template.
// - Content: A list of strings, where each string represents represents the path to a .html file you want executed.
// - Name: A string that indicates the name of the template you want executed. Use "" for no template execution
type FileStatic struct {
	Templates []string
	Content   interface{}
	Name      string
}
type FileStaticFunc func() FileStatic

// The GoX Router will render both the body or full page and serve based on state
// - Templates: A list of strings, where each string represents represents the path to a .html file you want executed.
// - Content: A struct containing the data you want executed in your template.
type FileDynamic struct {
	Templates []string
	Content   interface{}
}
type FileDynamicFunc func() FileDynamic

// The GoX Router will render and serve the full page
// - Value: A struct containing the data you want executed in your template.
// - StrArr: A list of strings, where each string represents represents the path to a .html file you want executed.
// - Str: A string that indicates the template you want executed. Use "" for no template execution
type TemplateStatic struct {
	Template *template.Template
	Content  interface{}
	Name     string
}
type TemplateStaticFunc func() TemplateStatic

// The GoX Router will render both the body or full page and serve based on state
// - Template: A Template object you used
// - Content: A struct containing the data you want executed in your template.
type TemplateDynamic struct {
	Template *template.Template
	Content  interface{}
}
type TemplateDynamicFunc func() TemplateDynamic
