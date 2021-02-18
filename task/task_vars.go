// Keep data about a task to be tested and interface to run all task's tests
package task

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/influxdata/kapacitor/client/v1"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// NewTaskVars constructor
func NewTaskVars(n string, p string) (*client.TaskVars, error) {
	fileVars := client.TaskVars{}
	if p != "" {
		if !strings.HasSuffix(p, "/") {
			p = p + "/"
		}
		f := p + n
		file, err := os.Open(f)
		defer file.Close()
		if err != nil {
			return nil, fmt.Errorf("failed to open file %v: %v", f, err)
		}

		data, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, fmt.Errorf("failed to read file %v: %v", f, err)
		}

		fn := file.Name()
		id := strings.TrimSuffix(filepath.Base(fn), filepath.Ext(fn))

		fileVars := client.TaskVars{}
		switch ext := path.Ext(f); ext {
		case ".yaml", ".yml":
			if err := yaml.Unmarshal(data, &fileVars); err != nil {
				return nil, errors.Wrapf(err, "failed to unmarshal yaml task vars file %q", f)
			}
		case ".json":
			if err := json.Unmarshal(data, &fileVars); err != nil {
				return nil, errors.Wrapf(err, "failed to unmarshal json task vars file %q", f)
			}
		default:
			return nil, errors.New("bad file extension. Must be YAML or JSON")
		}
		fileVars.ID = id
	}
	return &fileVars, nil
}
