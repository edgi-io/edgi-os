package module

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/edgi-io/edgi-os/pkg/config"
	"github.com/paultag/go-modprobe"
	"github.com/sirupsen/logrus"
)

const (
	procModulesFile = "/proc/modules"
)

func LoadModules(cfg *config.CloudConfig) error {
	loaded := map[string]bool{}
	f, err := os.Open(procModulesFile)
	if err != nil {
		return err
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		loaded[strings.SplitN(sc.Text(), " ", 2)[0]] = true
	}
	modules := cfg.EDGI.Modules
	for _, m := range modules {
		if loaded[m] {
			continue
		}
		params := strings.SplitN(m, " ", -1)
		logrus.Debugf("module %s with parameters [%s] is loading", m, params)
		if err := modprobe.Load(params[0], strings.Join(params[1:], " ")); err != nil {
			return fmt.Errorf("could not load module %s with parameters [%s], err %v", m, params, err)
		}
		logrus.Debugf("module %s is loaded", m)
	}
	return sc.Err()
}
