package mode

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"edgi.io/cmd/edgi/pkg/system"
)

func Get(prefix ...string) (string, error) {
	bytes, err := ioutil.ReadFile(filepath.Join(filepath.Join(prefix...), system.StatePath("mode")))
	if os.IsNotExist(err) {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(bytes)), nil
}
