package utils

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type PageData struct {
	Content   interface{}
	Templates []string
}

type RenderCustomFunc func() error
type RenderCustom struct {
	Handler RenderCustomFunc
}

// The GoX Router will render and serve the full page
// - Value: A struct containing the data you want executed in your template.
// - StrArr: A list of strings, where each string represents represents the path to a .html file you want executed.
// - Str: A string that indicates the template you want executed. Use "" for no template execution
type RenderFileStatic struct {
	Value  interface{}
	StrArr []string
	Str    string
}
type RenderFileStaticFunc func() RenderFileStatic

// The GoX Router will render both the body or full page and serve based on state
// - Value: A struct containing the data you want executed in your template.
// - StrArr: A list of strings, where each string represents represents the path to a .html file you want executed.
// - Str: A string that indicates the template you want executed. Use "" for no template execution
type RenderFileDynamic struct {
	Value  interface{}
	StrArr []string
	Str    string
}
type RenderFileDynamicFunc func() RenderFileDynamic

// The GoX Router will render and serve the full page
// - Value: A struct containing the data you want executed in your template.
// - StrArr: A list of strings, where each string represents represents the path to a .html file you want executed.
// - Str: A string that indicates the template you want executed. Use "" for no template execution
type RenderTemplateStatic struct {
	Value interface{}
	Tmpl  *template.Template
	Str   string
}
type RenderTemplateStaticFunc func() RenderTemplateStatic

// The GoX Router will render both the body or full page and serve based on state
// - Value: A struct containing the data you want executed in your template.
// - StrArr: A list of strings, where each string represents represents the path to a .html file you want executed.
// - Str: A string that indicates the template you want executed. Use "" for no template execution
type RenderTemplateDynamic struct {
	Value interface{}
	Tmpl  *template.Template
	Str   string
}
type RenderTemplateDynamicFunc func() RenderTemplateDynamic

type DataReturnType struct {
	PageData
	Error error
}

type DataReturnFunc func(w http.ResponseWriter, r *http.Request) DataReturnType

func RenderFile[T any](filePath string, outputDir string, templates []string, v T, templateExec string) error {
	tmpl := template.Must(template.ParseFiles(templates...))
	return RenderTemplate(filePath, outputDir, tmpl, v, templateExec)
}

func RenderTemplate[T any](filePath string, outputDir string, tmpl *template.Template, v T, templateExec string) error {
	file, err := CreateFile(filePath, outputDir)
	if err != nil {
		return err
	}

	err = WriteToFile(templateExec, file, tmpl, v)
	return err
}

func CreateFile(filePath string, outputDir string) (*os.File, error) {
	dir := filepath.Dir(outputDir + filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}
	file, err := os.Create(filepath.Join(outputDir, filePath))
	if err != nil {
		return nil, err
	}
	return file, nil
}

func WriteToFile[T any](templateExec string, file *os.File, tmpl *template.Template, v T) error {
	defer file.Close()

	var err error
	if templateExec == "" {
		err = tmpl.Execute(file, v)
	} else {
		err = tmpl.ExecuteTemplate(file, templateExec, v)
	}
	return err
}

func HandleGeneric[T any](templates []string, v T, w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles(templates...))
	var err error

	if IsHtmxRequest(r) {
		err = tmpl.ExecuteTemplate(w, "body", v)
	} else {
		err = tmpl.Execute(w, v)
	}

	if err != nil {
		log.Fatalf("execution: %s", err)
	}
}

func IsHtmxRequest(r *http.Request) bool {
	return r.Header.Get("HX-Request") == "true"
}

func GetHtmxRequestURL(r *http.Request) string {
	return r.Header.Get("HX-Current-URL")
}

func GenerateETag(content string) string {
	hash := md5.Sum([]byte(content))
	return fmt.Sprintf("%x", hash)
}

// StringSet represents a collection of unique strings.
type StringSet map[string]struct{}

// New creates a new StringSet.
func NewStringSet() StringSet {
	return make(StringSet)
}

// Add inserts the item into the set.
func (s StringSet) Add(item string) {
	s[item] = struct{}{}
}

// Remove deletes the item from the set.
func (s StringSet) Remove(item string) {
	delete(s, item)
}

// Contains checks if the item is in the set.
func (s StringSet) Contains(item string) bool {
	_, exists := s[item]
	return exists
}

// Elements returns the elements of the set as a slice of strings.
func (s StringSet) Elements() []string {
	elements := make([]string, 0, len(s))
	for key := range s {
		elements = append(elements, key)
	}
	return elements
}

// Join concatenates the elements of the set using the provided separator.
func (s StringSet) Join(separator string) string {
	return strings.Join(s.Elements(), separator)
}
