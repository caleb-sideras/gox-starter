package global

import (
	"bytes"
	"github.com/yuin/goldmark"
	"html/template"
	"io/ioutil"
	"strings"
)

func ConvertMarkdownToHTML(markdown []byte) (*template.Template, error) {
	var buf bytes.Buffer
	if err := goldmark.Convert(markdown, &buf); err != nil {
		return nil, err
	}
	tmpl, err := template.New("markdown").Parse(buf.String())
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

func MarkdownToHTML(path string) (*template.Template, error) {
	markdown, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	htmlTemplate, err := ConvertMarkdownToHTML(markdown)
	if err != nil {
		return nil, err
	}
	return htmlTemplate, nil
}

func MapKeysToSlice(m map[string]bool) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
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
