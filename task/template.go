// Keep data about a task to be tested and interface to run all task's tests
package task

import (
	"io/ioutil"
	"strings"
)

// FS configurations, namely path where TICKscripts templates are located
type Template struct {
	Name   string
	Path   string
	Script string
}

// Template constructor
func NewTemplate(n string, p string) (*Template, error) {
	template := Template{
		Name: n,
		Path: p}

	if p != "" {
		if !strings.HasSuffix(p, "/") {
			p = p + "/"
		}

		s, err := ioutil.ReadFile(p + n)
		if err != nil {
			return nil, err
		}
		template.Script = string(s[:])
	}
	return &template, nil
}
